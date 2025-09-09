package main

import (
	"entdemo/ent"
	"entdemo/internal/infrastructure"
	"fmt"
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

	host := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	usr := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPass := os.Getenv("POSTGRES_PASSWORD")

	connstr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, dbPort, usr, dbName, dbPass)

    client, err := ent.Open("postgres",connstr)

	if err != nil {
		panic(err.Error())
	}

	handlers := infrastructure.NewHandler(client)
	
	router.POST("/user", handlers.HandleCreateUser)
	router.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080
}