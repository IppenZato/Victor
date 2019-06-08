package appGuard

import (
	"os"
	"github.com/Viktor19931/books_api/utils"
)

var Engine = new_guard() // Maybe only 1 instance in system

type guard struct {
	name	string			// name of Engine
	started bool			// flg is started
	signals	[]os.Signal		// list of signal
	ch chan os.Signal		// channel of nahled signals
	ol 		[]interface{}   // list of guardian objects
}

func new_guard() *guard {
	this       := new(guard)
	this.name 	= utils.GetPackageName(this)
	this.ch 	= make(chan os.Signal, 1)
	this.ol 	=  nil
	return this
}

type stop_Accessor interface {
	Stop()
}

