package models

import (
	"github.com/Thaaaii/TODO-List/database"
	"github.com/Thaaaii/TODO-List/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func LoginCheck(username, password string) (string, error) {
	hashedPassword, err := SelectPassword(username)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(username)

	if err != nil {
		return "", err
	}
	return token, nil
}

func InsertUserIntoTable(username, password string) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO Users (name, password) VALUES (?, ?)", username, password)

	if err != nil {
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

func SelectPassword(username string) (string, error) {
	result := database.DB.QueryRow("SELECT password FROM Users WHERE name = ?", username)

	var password string
	err := result.Scan(&password)

	if err != nil {
		return "", err
	}

	return password, nil
}
