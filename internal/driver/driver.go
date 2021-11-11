package driver

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

)

func DBConnect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err	
	}

	log.Println("Connected to db!")

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Pinged db!")

	return db, nil
}