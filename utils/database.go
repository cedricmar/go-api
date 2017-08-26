package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// DBConnect returns db connection
func DBConnect() *sqlx.DB {
	if db != nil {
		return db
	}
	return sqlx.MustConnect("mysql", "admin:password@tcp(database:3306)/api")
}
