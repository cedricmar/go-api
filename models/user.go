package models

import "github.com/weebagency/go-api/utils"

type User struct {
	Id        int
	LastName  string `json:"last_name" db:"last_name"`
	FirstName string `json:"first_name" db:"first_name"`
	Address   string `json:"address"`
	City      string `json:city`
}

func GetUsers() []User {

	// Get DB
	db := utils.DBConnect()
	defer db.Close()

	// Get Users
	users := []User{}
	db.Select(&users, "SELECT * FROM user")
	return users
}
