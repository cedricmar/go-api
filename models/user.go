package models

import "github.com/weebagency/go-api/utils"

type User struct {
	Id        int
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Address   string `json:"address"`
	City      string `json:city`
}

func GetUsers() []*User {
	// Get Users
	db := utils.DBConnect()
	rows, err := db.Query("SELECT * FROM user")
	utils.LogFatal(err)
	// Map DB data to User type
	var users []*User
	for rows.Next() {
		u := new(User)
		err := rows.Scan(&u.Id, &u.LastName, &u.FirstName, &u.Address, &u.City)
		utils.LogFatal(err)

		users = append(users, u)
	}
	rows.Close()
	return users
}
