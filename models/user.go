package models

import (
	"github.com/Thaaaii/TODO-List/database"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func InsertUserIntoTable(username, password string) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO Users (name, password) VALUES (?, ?)", username, password)

	if err != nil {
		log.Fatal(err, username)
		return -1, err
	}

	return result.LastInsertId()
}

func SelectUserID(username string) (int64, error) {
	result := database.DB.QueryRow("SELECT id FROM Users WHERE name = ?", username)

	var id int64
	err := result.Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}
