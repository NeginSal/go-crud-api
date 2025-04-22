package main

import (
	"crud-api-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.UserRoutes(router)

	router.Run("localhost:8080")
}