package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitServer() {
	router := gin.Default()
	router.GET("user/tasks", getTasks)

	router.Run("localhost:8080")
}

func getTasks(ctx *gin.Context) {
	//user := ctx.Param(":user")
	ctx.IndentedJSON(http.StatusOK, SelectUserTasks(1))
}

func postTasks(ctx *gin.Context) {

}
