package migrate

import (
	"github.com/weebagency/go-api/utils"
)

// InitSchema is a migration function
func InitSchema() {
	up()
}

func up() {
	db := utils.DBConnect()
	defer db.Close()
	tx := db.MustBegin()
	// CREATE USER TABLE
	tx.MustExec("DROP TABLE user")
	tx.MustExec(`CREATE TABLE user (
		id INT(11) NOT NULL auto_increment,
		last_name VARCHAR(255) NOT NULL,
		first_name VARCHAR(255),
		email VARCHAR(255),
		PRIMARY KEY (id)
	)`)
	tx.MustExec("TRUNCATE TABLE user")
	tx.Commit()
}

func down() {
	db := utils.DBConnect()
	defer db.Close()
	tx := db.MustBegin()
	// DROP USER TABLE
	tx.MustExec("DROP TABLE IF EXISTS user")
	tx.Commit()
}
