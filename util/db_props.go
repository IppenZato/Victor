package util

import (
	"database/sql"
	"fmt"
)

// DbDriver
const DbDriver = "mysql"

// User
const User = "root"

// Pass
const Password = "11111111"

// DbName
const DbName = "db_go1"

// book table
const TableBook = "book"

// DataSourceName ...
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8",
	User, Password, DbName)

// DB
var DB *sql.DB

// RowCustomer ...
var RowCustomer *sql.Rows
