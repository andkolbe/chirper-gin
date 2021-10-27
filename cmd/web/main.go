package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine
var db *sql.DB

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	
	initializeRoutes()


	// Get a database handle
	var err error 
	db, err = sql.Open("mysql", dns)
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