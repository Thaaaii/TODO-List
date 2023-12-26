package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB = nil

// InitDatabase initializes the global database instance for further queries and requests.
func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./database/db.sqlite")

	if err != nil {
		log.Fatal(err)
	}
}

// CreateTables creates all the tables that are necessary for the program
func CreateTables() {
	createTableUsers()
	createTableTasks()
	createTableCategories()
}

// createTableUsers creates the scheme for the users table
func createTableUsers() {
	_, err := DB.Exec(`
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

// createTableTasks creates the scheme for the tasks table
func createTableTasks() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS Tasks (
			id INTEGER  PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			isDone BOOLEAN NOT NULL CHECK (isDone in (0, 1)),
		    sequence INTEGER DEFAULT 0,
			user_id INTEGER NOT NULL,
		    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

// createTableCategories creates the scheme for the categories table, which is needed for the tasks table
func createTableCategories() {
	_, err := DB.Exec(`
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
