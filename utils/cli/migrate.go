package cli

import (
	"log"

	"github.com/cedricmar/go-api/models"
	"github.com/cedricmar/go-api/utils"
)

var jane = &models.User{
	FirstName: "Jane",
	LastName:  "Citizen",
	Email:     "jane.citzen@example.com",
}

func Seed() {
	db := utils.DBConnect()
	defer db.Close()

	log.Println("--- Seeding ---")

	tx := db.MustBegin()
	tx.MustExec("TRUNCATE TABLE user")
	tx.MustExec("INSERT INTO user (first_name, last_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO user (first_name, last_name, email) VALUES (?, ?, ?)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.NamedExec("INSERT INTO user (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", jane)
	tx.Commit()

	log.Println("--- Seeding - Done ---")
}
