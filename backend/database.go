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

func InsertCategoriesIntoTable(categories []string, taskID int64) error {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	for _, label := range categories {
		_, err = db.Exec("INSERT INTO Categories (label, task_id) VALUES (?, ?)", label, taskID)

		if err != nil {
			log.Fatal(err)
		}
	}
	return err
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

	result, err := db.Query("SELECT Tasks.id, title, description, isDone FROM Tasks WHERE user_id = ?", userID)

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

		task.Categories = append(task.Categories, SelectTaskCategories(int64(task.ID))...)

		tasks = append(tasks, task)
	}

	return tasks
}

func SelectTaskCategories(taskID int64) []string {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Query("SELECT label FROM Categories WHERE task_id = ?", taskID)

	defer result.Close()

	if err != nil {
		log.Fatal(err)
	}

	categories := make([]string, 0)

	for result.Next() {
		var label string
		err := result.Scan(&label)

		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, label)
	}

	return categories
}

func UpdateTaskCategories(taskID int64, categories []string) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Print(err)
	}

	_, err = db.Exec(`
		DELETE FROM Categories 
		WHERE task_id = ?`,
		taskID,
	)

	if err != nil {
		log.Print(err)
	}

	err = InsertCategoriesIntoTable(categories, taskID)

	if err != nil {
		log.Print(err)
	}

}

func UpdateUserTask(taskID int64, title, description string, isDone bool) {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		UPDATE Tasks 
		SET title = ?, description = ?, isDone = ? 
		WHERE id = ?`,
		title, description, isDone, taskID,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateTables() {
	createTableUsers()
	createTableTasks()
	createTableCategories()
}

func createTableUsers() {
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

func createTableTasks() {
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

func createTableCategories() {
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
