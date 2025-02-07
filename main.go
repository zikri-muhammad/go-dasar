package main

import (
	"fmt"
	"net/http"
	"strconv"

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

func Move(position int, roll int) int {
	// your code here
	return position + roll*2
}

func ToCsvText(array [][]int) string {
	// your code here
	var result string
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			result += strconv.Itoa(array[i][j])
			if j < len(array[i])-1 {
				result += ","
			}
		}
		if i < len(array)-1 {
			result += "\n"
		}
	}
	return result
}

func main() {
	fmt.Println(Move(0, 4))
	fmt.Println(ToCsvText([][]int{
		{0, 1, 2, 3, 45},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34}}))
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// // New endpoint to get user by ID
	// r.GET("/user/:id", getUser)

	// r.Run(":8080")
}
