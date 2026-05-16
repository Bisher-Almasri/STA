// @title           Second Tech API
// @version         1.0
// @description     This is a replacement of FTC Events api as it always goes down and inconsistent with its changes

// @contact.name   Bisher Almasri
// @contact.email  bisherk.almasri@gmail.com

// @license.name  CC BY NC 4.0
// @license.url   https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

package main

import (
	"STA/handlers"
	"net/http"
	"time"
)

var client = http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	mux := http.NewServeMux()


	mux.Handle(
	"/docs/",
	http.StripPrefix(
		"/docs/",
		http.FileServer(http.Dir("./docs")),
	))

		mux.Handle(
	"/static/",
	http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./static")),
	))

	mux.HandleFunc("/api/v1/teams/", handlers.GetTeam(client))
	mux.HandleFunc("/api/v1/events/", handlers.GetEvent(client))

	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
