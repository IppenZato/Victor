package db

import (
	"fmt"
	"github.com/Viktor19931/books_api/utils"
)

var name = utils.GetPackageName(&_dbServer{})
var DbServer = &_dbServer{Alias: "default"}
var al *alias = newAlias()

type _dbServer struct {
	Alias    string `orm:"column(alias)"    json:"alias"    desc:"DB alias"`
	Driver   string `orm:"column(driver)"   json:"driver"   desc:"DB driver"     oneof:"postgres"`
	Host     string `orm:"column(host)"     json:"host"     desc:"DB host"`
	Port     int    `orm:"column(port)"     json:"port"     desc:"DB port"`
	DbName   string `orm:"column(dbname)"   json:"dbname"   desc:"DB name"`
	Schema   string `orm:"column(schema)"   json:"schema"   desc:"DB shema name"`
	User     string `orm:"column(user)"     json:"user"     desc:"DB user"`
	Password string `orm:"column(password)" json:"password" desc:"DB password"`
}

func (this *_dbServer) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		this.User, this.Password, this.Host, this.Port, this.DbName)
}
