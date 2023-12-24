package backend

import (
	"github.com/Thaaaii/TODO-List/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()

	router.LoadHTMLFiles("frontend/todo/index.html", "frontend/login/login.html")
	router.Static("img", "./img")
	router.Static("/static", "./frontend/todo")
	router.Static("/login-static", "./frontend/login")
	router.GET("/todo-list/:user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "TODO-Liste",
		})
	})
	router.GET("login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})

	router.GET("/:user/tasks", getTasks)
	router.POST("/:user/tasks", postTasks)
	router.PATCH("/todo-list/:user/tasks/:taskID", patchTask)
	router.DELETE("/todo-list/:user/tasks/:taskID", deleteTask)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}

func getTasks(ctx *gin.Context) {
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

func postTasks(ctx *gin.Context) {
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

func patchTask(ctx *gin.Context) {
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

func deleteTask(ctx *gin.Context) {
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
