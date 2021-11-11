package main

import (
	"log"
	"os"

	"github.com/andkolbe/chirper-gin/internal/driver"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

func main() {
	LoadEnv()
	URL := os.Getenv("URL")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// connect to db
	_, err := driver.DBConnect(URL)
	if err != nil {
		log.Fatal(err)
	}

	r.Run("127.0.0.1:8080")
}
