package server

import (
	"fmt"
	"iplogger/db"
	"net/http"
)

type HttpServer struct {
	db db.DB
	logger Logging
}

func CreateServer() Server {
	db, err := db.New()
	if err != nil {
		panic("Can't connect to SQLite DB")
	}
	return HttpServer{
		db: db,
		logger: GetLogger(),
	}
}

func (httpServer HttpServer) Run()  {
	http.HandleFunc("/urls", func(w http.ResponseWriter, r *http.Request) {
		httpServer.logger.Log(fmt.Sprintf("%s - %s", r.RequestURI, r.UserAgent()))
	})
	http.ListenAndServe(":8080", nil)
}
