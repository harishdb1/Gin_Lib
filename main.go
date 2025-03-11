package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a route for the root path
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World! This is my first Program with Gin_Lib",
		})
	})

	// Define a route for a GET request to /greet/:name
	r.GET("/greet/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
}
