package routes

import (
	"crud-api-go/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUserByID)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}
