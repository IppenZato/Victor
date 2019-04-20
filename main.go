package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	utl "books-list/util"

	_ "github.com/go-sql-driver/mysql" // let work sql
)

func init() {
	rand.Seed(time.Now().UnixNano())

	maxCPU := runtime.NumCPU()

	utl.CPUUsed = 4 // use 4 CPU
	runtime.GOMAXPROCS(utl.CPUUsed)

	fmt.Printf("\n=========================================\n")
	fmt.Printf("= Number of CPUs (Total=%d - Used=%d)", maxCPU, utl.CPUUsed)
	fmt.Printf("\n=========================================\n\n")
}

func main() {
	fmt.Printf("Opening the %s ...\n\n", utl.DbName)

	var err error
	utl.DB, err = sql.Open(utl.DbDriver, utl.DataSourceName)

	defer utl.DB.Close()

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success !!\n")
	}

	utl.BookService()
}
