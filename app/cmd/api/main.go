package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("API_PORT")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080
}
