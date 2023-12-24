package controller

import (
	"github.com/Thaaaii/TODO-List/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostUser(ctx *gin.Context) {
	var newUser models.User

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := models.InsertUserIntoTable(newUser.Name, newUser.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Task could not be inserted",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func Login(ctx *gin.Context) {

}

func GetTasks(ctx *gin.Context) {
	var userTasks []models.Task
	user := ctx.Param("user")
	userID, err := models.SelectUserID(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "User could not be found",
		})
		return
	}

	userTasks, err = models.SelectUserTasks(userID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Tasks of the user could not be found",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, userTasks)
}

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
