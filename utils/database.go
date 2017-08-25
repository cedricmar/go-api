package services

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBConnect() *sql.DB {
	if db != nil {
		return db
	}
	conn, err := sql.Open("mysql", "admin:130779@/api")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	db = conn
	return conn
}
