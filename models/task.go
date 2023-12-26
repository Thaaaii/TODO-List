package models

import (
	"github.com/Thaaaii/TODO-List/database"
)

// Task defines a structs to contain all relevant details about tasks
type Task struct {
	ID             int      `json:"id"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Categories     []string `json:"categories"`
	IsDone         bool     `json:"is_done"`
	SequenceNumber int      `json:"sequenceNumber"`
}

// InsertTaskIntoTable defines a SQL statement to insert tasks into the table
func InsertTaskIntoTable(title, description string, isDone bool, userID int64) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO Tasks (title, description, isDone, user_id) VALUES (?, ?, ?, ?)", title, description, isDone, userID)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

// InsertCategoriesIntoTable defines a SQL statement to insert categories into the table
// and also links the categories to the corresponding task through a foreign key
func InsertCategoriesIntoTable(categories []string, taskID int64) error {
	for _, label := range categories {
		_, err := database.DB.Exec("INSERT INTO Categories (label, task_id) VALUES (?, ?)", label, taskID)

		if err != nil {
			return err
		}
	}
	return nil
}

// SelectUserTasks queries all tasks of a specific user and returns them to the caller
func SelectUserTasks(userID int64) ([]Task, error) {
	result, err := database.DB.Query("SELECT Tasks.id, title, description, isDone, sequence FROM Tasks WHERE user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0)

	for result.Next() {
		var task Task
		err := result.Scan(&task.ID, &task.Title, &task.Description, &task.IsDone, &task.SequenceNumber)

		if err != nil {
			return nil, err
		}

		categories, err := SelectTaskCategories(int64(task.ID))

		if err != nil {
			return nil, err
		}

		task.Categories = append(task.Categories, categories...)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// SelectTaskCategories queries all categories of a specific task and returns them to the caller
func SelectTaskCategories(taskID int64) ([]string, error) {
	result, err := database.DB.Query("SELECT label FROM Categories WHERE task_id = ?", taskID)

	if err != nil {
		return nil, err
	}

	categories := make([]string, 0)

	for result.Next() {
		var label string
		err := result.Scan(&label)

		if err != nil {
			return nil, err
		}

		categories = append(categories, label)
	}
	return categories, nil
}

// UpdateTaskCategories deletes all current categories of a task and inserts the updated categories into the table
func UpdateTaskCategories(taskID int64, categories []string) error {
	_, err := database.DB.Exec(`
		DELETE FROM Categories 
		WHERE task_id = ?`,
		taskID,
	)

	if err != nil {
		return err
	}

	err = InsertCategoriesIntoTable(categories, taskID)

	if err != nil {
		return err
	}
	return nil
}

// UpdateUserTask updates a task of a user
func UpdateUserTask(taskID int64, title, description string, isDone bool) error {
	_, err := database.DB.Exec(`
		UPDATE Tasks 
		SET title = ?, description = ?, isDone = ? 
		WHERE id = ?`,
		title, description, isDone, taskID,
	)

	if err != nil {
		return err
	}
	return nil
}

// UpdateUserTaskOrder updates the sequence number of a user task
func UpdateUserTaskOrder(taskID, sequenceNumber int64) error {
	_, err := database.DB.Exec(`
		Update Tasks
		Set sequence = ?
		WHERE id = ?`,
		sequenceNumber, taskID,
	)

	if err != nil {
		return err
	}
	return nil
}

// DeleteUserTask deletes a user task selected by task ID
func DeleteUserTask(taskID int64) error {
	_, err := database.DB.Exec(`
		DELETE FROM Tasks
		WHERE id = ?`,
		taskID,
	)

	if err != nil {
		return err
	}
	return nil
}
