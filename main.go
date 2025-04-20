package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var users = []User{
	{ID: 1, Name: "mario", Email: "mario@gmail.com"},
	{ID: 2, Name: "luigi", Email: "luigi@gmail.com"},
}

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run("localhost:8080")
}
