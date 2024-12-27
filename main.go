package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// getUser handles GET requests to fetch a user by ID
func getUser(c *gin.Context) {
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	c.JSON(http.StatusOK, user)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// New endpoint to get user by ID
	r.GET("/user/:id", getUser)

	r.Run(":8080")
}
