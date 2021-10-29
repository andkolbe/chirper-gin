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

// create a custom Env struct which holds a connection pool
// all the dependencies for our handlers are explicitly defined in one place
type Env struct {
	users interface {
		GetAllUsers() ([]models.User, error)
	}
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

func main() {
	LoadEnv()
	
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)

	// Get a database handle
	db, err := sql.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
	}

	// because sql.Open doesn't actually check a connection, we also call DB.Ping() to make sure that everything works OK on startup
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to DB!")

	// create an instance of Env containing the connection pool
	env := &Env{
		users: models.UserModel{DB: db},
	}

	router := gin.Default()

	router.GET("/users", env.UsersIndex)

	router.Run()	
}


func (env *Env) UsersIndex(c *gin.Context) {
	users, err := env.users.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		return
	}

	for _, user := range users {
		fmt.Printf("%q, %s, %s, %s", user.ID, user.Name, user.Email, user.Password)
	}
}
