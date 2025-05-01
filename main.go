package main

import (
	"crud-api-go/config"
	"crud-api-go/routes"
	"github.com/gin-gonic/gin"

	// Swagger
	_ "crud-api-go/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title CRUD API with Go and Gin
// @version 1.0
// @description This is a sample CRUD API server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /

func main() {
	config.ConnectDB()

	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(router)

	router.Run("localhost:8080")

}
