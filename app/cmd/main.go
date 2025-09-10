package main

import (
	"entdemo/ent"
	"entdemo/internal/infrastructure"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	port := os.Getenv("API_PORT")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	dbConn := infrastructure.NewDbConn()

	client, err := ent.Open(dbConn.Driver, dbConn.Connstr)

	if err != nil {
		panic(err.Error())
	}

	handlers := infrastructure.NewHandler(client)

	router.POST("/user", handlers.HandleCreateUser)
	router.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080
}
