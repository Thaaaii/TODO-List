package routes

import (
	"github.com/Thaaaii/TODO-List/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter() {
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

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	router.GET("/:user/tasks", controller.GetTasks)
	router.POST("/:user/tasks", controller.PostTasks)
	router.PATCH("/todo-list/:user/tasks/:taskID", controller.PatchTask)
	router.DELETE("/todo-list/:user/tasks/:taskID", controller.DeleteTask)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
