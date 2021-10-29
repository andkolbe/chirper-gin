package models

import (
	"database/sql"
	"time"
)

// we can only use string and int safely because we set NOT NULL constraints on all of the columns on the table
type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	// we fetch a result set from the books table using the DB.Query() method and assign it to a rows variable
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	// defer ensures the result set is properly closed before the parent function returns
	defer rows.Close()

	var users []User

	// use rows.Next() to iterate through the rows in the result set. This preps every row to be acted on by the rows.Scan() method
	for rows.Next() {
		var user User
		// use rows.Scan() method to copy the values from each field in the row to a new User object that we created 
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Created_At)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	// it's important to call rows.Err(). Don't assume that we completed a successful iteration over the whole result set 
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}