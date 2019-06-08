package server

import (
	"fmt"
)

var server = newServer()

func newServer() *_server {
	this := &_server{}
	this.Port = 8000 // default port
	return this
}

type _server struct {
	Port        int
	started     bool
}

func (this *_server) Addr() string {
	return fmt.Sprintf(":%v", this.Port)
}
