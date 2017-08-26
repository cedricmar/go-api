package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// DBConnect returns db connection
func DBConnect() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "admin:password@tcp(database:3306)/api")
	if err != nil {
		panic(err)
	}
	return db
}
