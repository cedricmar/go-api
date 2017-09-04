package models

import "github.com/weebagency/go-api/utils"

type User struct {
	Id        int
	LastName  string `json:"last_name" db:"last_name"`
	FirstName string `json:"first_name" db:"first_name"`
	Email     string `json:"email" db:"email"`
}

func GetUsers() []User {

	// Get DB
	db := utils.DBConnect()
	defer db.Close()

	// Get Users
	users := []User{}
	err := db.Select(&users, "SELECT * FROM user")
	utils.LogFatal(err)
	return users
}
