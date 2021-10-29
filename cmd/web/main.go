package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/andkolbe/chirper-gin/internal/env"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine
var db *sql.DB

func main() {
	env.LoadEnv()

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	
	initializeRoutes()

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)

	// Get a database handle
	var err error 
	db, err = sql.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("connected to DB!")

	defer db.Close()

	router.Run()
}