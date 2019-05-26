package db

import (
	"github.com/Viktor19931/books_api/log"
	"github.com/Viktor19931/books_api/app"
	"time"
	"fmt"
	"github.com/Viktor19931/books_api/utils"
	"github.com/Viktor19931/books_api/db/sqlx"
)

const ini_section = "db"

type tDB struct { name string }
var  vDB = &tDB{name: utils.GetPackageName(&tDB{})}

func Start() *tDB {
	sec := app.Config.IniFile.Section(ini_section)

	DbServer.Driver   = sec.Key("Driver").MustString("mysql")
	DbServer.Host 	  = sec.Key("Host").MustString("localhost")
	DbServer.Port     = sec.Key("Port").MustInt(3306)
	DbServer.DbName   = sec.Key("DbName").MustString("db_go1")
	DbServer.User 	  = sec.Key("User").MustString("root")
	DbServer.Password = sec.Key("Password").MustString("11111111")
	app.Config.SaveLastIni()

	open() // try connect to database

	return vDB
}

func (t *tDB) Stop() {

}

func open() {
	log.Debug("connecting to: %s\n", DbServer.GetConnectionString())

	db, err := sqlx.Connect(DbServer.Driver, DbServer.GetConnectionString())
	if err != nil {
		panic(fmt.Sprintf("Error connection to %q:%q(user:%q)", DbServer.Driver, DbServer.DbName, DbServer.User))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("Error ping:%v", err))
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(0) // forever


	// fill alias data
	al.TZ         = time.UTC	// Default TimeZone
	al.Server     = DbServer
	al.Name       = DbServer.Alias
	al.DriverName = DbServer.Driver
	if dr, ok := drivers[DbServer.Driver]; ok {
		al.DriverType = dr
		//al.DbBaser    = dbBasers[dr]
	} else {
		panic(fmt.Errorf("driver name `%s` have not registered", DbServer.Driver))
	}
	al.MaxIdleConns = 100
	al.MaxOpenConns = 10
	al.DB = db

	log.Debug("connected to: %s\n", DbServer.GetConnectionString())
	return
}

