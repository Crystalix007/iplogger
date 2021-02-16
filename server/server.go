package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"iplogger/db"
	"net/http"
)

type HttpServer struct {
	db     db.DB
	logger Logging
}

// CreateServer is a factory function to create a Server instance
func CreateServer() Server {
	db, err := db.New()
	if err != nil {
		panic("Can't connect to SQLite DB")
	}
	return HttpServer{
		db:     db,
		logger: GetLogger(),
	}
}

func createHandler(server *HttpServer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		shortUrl := params["url"]
		server.GetLongUrl(shortUrl)
		server.logger.Log(fmt.Sprintf("Get long url %s", shortUrl))
		server.logger.Log(fmt.Sprintf("%s - %s", r.RequestURI, r.UserAgent()))
	}
}

func (httpServer HttpServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/urls/{url}", createHandler(&httpServer))
	router.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		httpServer.NewUrl(params["url"])
	})
	httpServer.db.Migrate()
	http.ListenAndServe(":8080", router)
}

func (httpServer HttpServer) GetLongUrl(shortUrl string) string {
	db := &httpServer.db
	val, err := db.GetFullUrlFromShort(shortUrl)
	if err == nil {
		return val
	} else {
		return "nofound"
	}
}

func (httpServer HttpServer) NewUrl(longUrl string) string {
	db := &httpServer.db
	return db.CreateUrl(longUrl)
}

func (httpServer HttpServer) Migrate() {
	httpServer.db.Migrate()
}
