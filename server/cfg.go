package server

import "github.com/Viktor19931/books_api/app"

func initServer() {
	sec := app.Config.IniFile.Section(ini_section)
	server.Port = sec.Key("Port").MustInt(8000)
	app.Config.SaveLastIni()
}
