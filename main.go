package main

import (
	"crud-api-go/config"
	"crud-api-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	router := gin.Default()

	routes.UserRoutes(router)

	router.Run("localhost:8080")

}
