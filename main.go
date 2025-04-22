package main

import (
	"crud-api-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
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

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if id == strconv.Itoa(user.ID) {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found."})
}

func createUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User

	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}

	for i, user := range users {
		if id == strconv.Itoa(user.ID) {
			users[i] = updatedUser
			c.IndentedJSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User Not Found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if id == strconv.Itoa(user.ID) {
			users = append(users[:i], users[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
