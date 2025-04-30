package routes

import (
	"crud-api-go/controller"
	"crud-api-go/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", controller.GetUsers)
		protected.GET("/users/:id", controller.GetUserByID)
		protected.POST("/users", controller.CreateUser)
		protected.PUT("/users/:id", controller.UpdateUser)
		protected.DELETE("/users/:id", controller.DeleteUser)
	}
}
