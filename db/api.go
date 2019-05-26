package db

import (
	_ "github.com/go-sql-driver/mysql" // let work sql
	_ "github.com/lib/pq"
	"database/sql"
	"context"
	"errors"
	"github.com/Viktor19931/books_api/db/sqlx"
)
var ErrNoConn = errors.New("database/sql: no connection")

func BeginTransaction() (*sql.Tx, error) {
	if al.DB == nil {
		return nil, ErrNoConn
	}
	return al.DB.Begin()
}

func Rollback(tx *sql.Tx) error {
	if nil == tx {
		return errors.New("Transaction is nil")
	}
	return tx.Rollback()
}

func Commit(tx *sql.Tx) error {
	if nil == tx {
		return errors.New("Transaction is nil")
	}
	return tx.Commit()
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	if al.DB == nil {
		return nil
	}
	return al.DB.QueryRow(query, args...)
}

func QueryRows(query string) (*sql.Rows, error) {
	if al.DB == nil {
		return nil, ErrNoConn
	}
	return al.DB.QueryContext(context.Background(), query)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	if al.DB == nil {
		return nil, ErrNoConn
	}
	return al.DB.Exec(query, args...)
}

func Instance() *sqlx.DB {
	return al.DB
}

