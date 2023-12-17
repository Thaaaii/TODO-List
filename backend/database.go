package backend

import (
	"database/sql"
	"log"
)

func InsertUserIntoTable(username, password string) (int64, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO Users (name, password) VALUES (?, ?)", username, password)

	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

func InsertTaskIntoTable(title, description string, isDone bool, userID int64) (int64, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO Tasks (title, description, isDone, user_id) VALUES (?, ?, ?, ?)", title, description, isDone, userID)

	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

func InsertCategoryIntoTable(label string, taskID int64) (int64, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO Categories (label, task_id) VALUES (?, ?)", label, taskID)

	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

func SelectUserID(username string) int64 {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	result := db.QueryRow("SELECT id FROM Users WHERE name = ?", username)

	var id int64
	err = result.Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func SelectUserTasks(userID int64) []Task {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Query("SELECT id, title, description, isDone FROM Tasks WHERE user_id = ?", userID)
	defer result.Close()

	if err != nil {
		log.Fatal(err)
	}

	tasks := make([]Task, 0)

	for result.Next() {
		var task Task
		err := result.Scan(&task.ID, &task.Title, &task.Description, &task.IsDone)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func CreateTableUsers() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
			id INTEGER  PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateTableTasks() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Tasks (
			id INTEGER  PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			isDone BOOLEAN NOT NULL CHECK (isDone in (0, 1)),
			user_id INTEGER NOT NULL,
		    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateTableCategories() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Categories (
			id INTEGER  PRIMARY KEY AUTOINCREMENT,
			label TEXT NOT NULL,
			task_id INTEGER NOT NULL,
		    FOREIGN KEY (task_id) REFERENCES Tasks(id) ON DELETE CASCADE
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}
