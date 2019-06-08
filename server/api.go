package server

import (
	"github.com/Viktor19931/books_api/log"
	"net/http"
)

const ini_section = "server"

func Start() *_server {
	if server.started {
		return server
	}

	initServer()
	InitRouters()

	log.Info("Server starting on %v", server.Addr())
	err := http.ListenAndServe(server.Addr(), nil) // задаем слушать порт
	if err != nil {
		log.Error("ListenAndServe: ", err)
	}
	server.started = true
	return server
}
