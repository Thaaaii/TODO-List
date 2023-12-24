package models

import "github.com/Thaaaii/TODO-List/database"

type Task struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	IsDone      bool     `json:"is_done"`
}

func InsertTaskIntoTable(title, description string, isDone bool, userID int64) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO Tasks (title, description, isDone, user_id) VALUES (?, ?, ?, ?)", title, description, isDone, userID)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func InsertCategoriesIntoTable(categories []string, taskID int64) error {
	for _, label := range categories {
		_, err := database.DB.Exec("INSERT INTO Categories (label, task_id) VALUES (?, ?)", label, taskID)

		if err != nil {
			return err
		}
	}
	return nil
}

func SelectUserTasks(userID int64) ([]Task, error) {
	result, err := database.DB.Query("SELECT Tasks.id, title, description, isDone FROM Tasks WHERE user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0)

	for result.Next() {
		var task Task
		err := result.Scan(&task.ID, &task.Title, &task.Description, &task.IsDone)

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
