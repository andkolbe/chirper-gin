package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var router *gin.Engine

var db *sql.DB

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
}

func GetAllUsers(c *gin.Context) {
	fmt.Println("Got here 0")
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Got here 1")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Created_At)
		if err != nil {
			fmt.Println("Got here 2")
			fmt.Println(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Got here 3")
		fmt.Println(err)
	}

	for _, user := range users {
		fmt.Printf("%q, %s, %s, %s", user.ID, user.Name, user.Email, user.Password)
	}

}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

func initializeRoutes() {
	router.GET("/users", GetAllUsers)
}

func main() {
	LoadEnv()
	
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

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to DB!")

	defer db.Close()

	router = gin.Default()

	initializeRoutes()

	router.Run()
	
}