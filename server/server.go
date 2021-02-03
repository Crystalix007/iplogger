package server

import (
	"fmt"
	"iplogger/db"
	"net/http"
	"strings"
	"time"
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

// Run starts the webserver
func (httpServer HttpServer) Run() {
	http.HandleFunc("/urls/", func(w http.ResponseWriter, r *http.Request) {
		shortUrl := strings.Replace(r.URL.Path, "/urls/", "", 1)

		// Log timestamp, remote address, cookies, referer, and user agent.
		// Add anything else you want to log below.
		httpServer.logger.Log(fmt.Sprintf("[url: %s] [%s] %s - %s - %s - %s",
			shortUrl,
			time.Now(),
			r.RemoteAddr,
			r.Cookies(),
			r.Referer(),
			r.UserAgent(),
		))

		

	})
	http.ListenAndServe(":8080", nil)
}
