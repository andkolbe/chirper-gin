package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andkolbe/chirper-gin/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

func initializeRoutes() {
	router.GET("/users", UsersIndex)
}

// Go handles the execution of this function for us
func init() {
	LoadEnv()
	
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)

	// Get a database handle
	var err error 
	models.DB, err = sql.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
	}

	// because sql.Open doesn't actually check a connection, we also call DB.Ping() to make sure that everything works OK on startup
	if err = models.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to DB!")
}

func main() {

	router = gin.Default()

	initializeRoutes()

	router.Run()
	
}


func UsersIndex(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		return
	}

	for _, user := range users {
		fmt.Printf("%q, %s, %s, %s", user.ID, user.Name, user.Email, user.Password)
	}
}