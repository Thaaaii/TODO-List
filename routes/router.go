package routes

import (
	"github.com/Thaaaii/TODO-List/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// InitRouter initializes the router and configures the specific routes
func InitRouter() {
	router := gin.Default()

	router.LoadHTMLFiles("frontend/index.html", "frontend/login.html")
	router.Static("/img", "./frontend/img")
	router.Static("/static", "./frontend/static")

	router.GET("login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)

	protected := router.Group("todo-list")
	protected.Use(controller.AuthenticationMiddleware())

	{
		protected.GET("/:user", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Todo-Liste",
			})
		})
		protected.GET("/:user/tasks", controller.GetTasks)
		protected.POST("/:user/tasks", controller.PostTasks)
		protected.PATCH("/:user/tasks/:taskID", controller.PatchTask)
		protected.PATCH("/:user/tasks/:taskID/order", controller.PatchTaskOrder)
		protected.DELETE("/:user/tasks/:taskID", controller.DeleteTask)
	}

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
