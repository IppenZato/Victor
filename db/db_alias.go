package db

import (
	"time"
	"github.com/Viktor19931/books_api/db/sqlx"
)

// the Database driver's Name
const (
	DBMySql = "mysql"
)

// DriverType database driver constant int.
type DriverType int

// Enum the Database driver
const (
	_       DriverType = iota // int enum type
	DRMySql                   // mysql
)

var drivers = map[string]DriverType{ DBMySql: DRMySql}

const wait_keepalive = 60 * time.Second

type alias struct {
	Name         string
	Server		 *_dbServer
	DriverName   string
	DriverType   DriverType
	MaxIdleConns int
	MaxOpenConns int
	DB           *sqlx.DB
	//DbBaser      dbBaser
	TZ           *time.Location
	//Engine       string
	active		 chan bool
}

func newAlias() *alias {
	al := new(alias)
	al.active = make(chan bool)
	return al
}
