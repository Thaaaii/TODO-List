package backend

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()

	router.LoadHTMLFiles("frontend/index.html")
	router.Static("img", "./img")
	router.Static("/static", "./frontend")
	router.GET("/todo-list/:user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "TODO-Liste",
		})
	})

	router.GET("/:user/tasks", getTasks)
	router.POST("/:user/tasks", postTasks)
	router.PATCH(":user/tasks/:taskID", patchTask)
	router.DELETE(":user/tasks/:taskID", deleteTask)

	router.Run("localhost:8080")
}

func getTasks(ctx *gin.Context) {
	user := ctx.Param("user")
	userID := SelectUserID(user)
	ctx.IndentedJSON(http.StatusOK, SelectUserTasks(userID))
}

func postTasks(ctx *gin.Context) {
	user := ctx.Param("user")
	var newTask Task

	if err := ctx.BindJSON(&newTask); err != nil {
		log.Fatal(err)
	}

	userID := SelectUserID(user)
	taskID, err := InsertTaskIntoTable(newTask.Title, newTask.Description, newTask.IsDone, userID)

	if err != nil {
		log.Fatal(err)
	}

	err = InsertCategoriesIntoTable(newTask.Categories, taskID)

	if err != nil {
		log.Fatal(err)
	}

	ctx.IndentedJSON(http.StatusCreated, newTask)
}

func patchTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("taskID"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	var updatedTask Task

	if err = ctx.BindJSON(&updatedTask); err != nil {
		log.Fatal(err)
	}

	UpdateUserTask(taskID, updatedTask.Title, updatedTask.Description, updatedTask.IsDone)
	UpdateTaskCategories(taskID, updatedTask.Categories)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ressource with ID " + ctx.Param("taskID") + " updated.",
	})
}

func deleteTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("taskID"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	DeleteUserTask(taskID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ressource deleted",
		"id":      taskID,
	})
}
