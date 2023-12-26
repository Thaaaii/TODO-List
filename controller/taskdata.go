package controller

import (
	"github.com/Thaaaii/TODO-List/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetTasks is a handler that processes GET requests. The handler queries all tasks of a specific user from
// the database and returns it to the specific route.
func GetTasks(ctx *gin.Context) {
	var userTasks []models.Task
	user := ctx.Param("user")
	userID, err := models.SelectUserID(user)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "User could not be found",
		})
		return
	}

	userTasks, err = models.SelectUserTasks(userID)

	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Tasks of the user could not be found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, userTasks)
}

// PostTasks is a handler that processes POST requests. The handler receives data from the frontend to create a new
// task entry in the database.
func PostTasks(ctx *gin.Context) {
	user := ctx.Param("user")
	var newTask models.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, err := models.SelectUserID(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User could not be found",
		})
		return
	}

	taskID, err := models.InsertTaskIntoTable(newTask.Title, newTask.Description, newTask.IsDone, userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Task could not be inserted",
		})
		return
	}

	newTask.ID = int(taskID)
	err = models.InsertCategoriesIntoTable(newTask.Categories, taskID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Categories of the task could not be inserted",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newTask)
}

// PatchTask is a handler that processes PATCH requests. The handler receives updates about changes in tasks
// to update the entry in the database.
func PatchTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("taskID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updatedTask models.Task

	if err = ctx.BindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.UpdateUserTask(taskID, updatedTask.Title, updatedTask.Description, updatedTask.IsDone)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Task to update could not be found",
		})
		return
	}

	err = models.UpdateTaskCategories(taskID, updatedTask.Categories)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Task to update categories could not be found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ressource with ID " + ctx.Param("taskID") + " updated.",
	})
}

// PatchTaskOrder is a handler that processes PATCH requests. The handler receives updates about changes in
// task order and updates the database accordingly.
func PatchTaskOrder(ctx *gin.Context) {
	type Sequence struct {
		SequenceNumber int `json:"sequenceNumber"`
	}

	taskID, err := strconv.ParseInt(ctx.Param("taskID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var sequence Sequence

	if err = ctx.BindJSON(&sequence); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.UpdateUserTaskOrder(taskID, int64(sequence.SequenceNumber))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Task to update could not be found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ressource with ID " + ctx.Param("taskID") + " updated.",
	})
}

// DeleteTask is a handler that processes DELETE requests. The handler deletes specific entries in the database.
func DeleteTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("taskID"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.DeleteUserTask(taskID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Task could not be deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ressource deleted",
		"id":      taskID,
	})
}
