package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBConnect() *sql.DB {
	conn, err := sql.Open("mysql", "admin:130779@tcp(database:3306)/api")
	if err != nil {
		panic(err)
	}
	db = conn
	return conn
}
