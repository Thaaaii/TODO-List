package backend

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()
	router.GET("/:user/tasks", getTasks)
	router.POST("/:user/tasks", postTasks)
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

	_, err = InsertCategoriesIntoTable(newTask.Categories, taskID)

	if err != nil {
		log.Fatal(err)
	}

	ctx.IndentedJSON(http.StatusCreated, newTask)
}
