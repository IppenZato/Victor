package main

import (
	"github.com/Viktor19931/books_api/log"
	"github.com/Viktor19931/books_api/app"
	"github.com/Viktor19931/books_api/appGuard"
	"github.com/Viktor19931/books_api/utils"
	"syscall"
	"github.com/Viktor19931/books_api/server"
	"github.com/Viktor19931/books_api/db"
)
func main() {
	defer log.AppRecover(main, nil)	// перехватчик аварийного завершения программы

	log.Info("starting %v...", utils.RunFile())

	////------------------------------------------------------------
	// перехватчик терминации проги - содержит список запускаемыемфх модулей
	// при выходе приложения запускает функцию Stop() всех модулей
	//------------------------------------------------------------
	gapp := appGuard.Start(syscall.SIGINT, syscall.SIGTERM)
	gapp.Add( app.Start() )     	// инициализация приложения
	gapp.Add( db.Start() )     		// инициализация базы данных
	gapp.Add( server.Start())  		// инициализация базы данных

	//gapp.WaitSignal()
}
