package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func InsertUserIntoTable(username, password string) (int64, error) {
	result, err := db.Exec("INSERT INTO Users (name, password) VALUES (?, ?)", username, password)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func SelectUserID(username string) (int64, error) {
	result := db.QueryRow("SELECT id FROM Users WHERE name = ?", username)

	var id int64
	err := result.Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}
