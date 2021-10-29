package models

import (
	"database/sql"
)

type Chirp struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	Location string `json:"location"`
}

var db *sql.DB

// get all chirps
// get one chirp
// add a new chirp
// update a chirp
// destroy a chirp

func ShowAllChirps() ([]Chirp, error) {
	// a chirp slice to hold data from returned rows
	var chirps []Chirp

	rows, err := db.Query("SELECT * FROM chirps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// loop through rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var chirp Chirp
		if err := rows.Scan(
			&chirp.ID,
			&chirp.Content,
			&chirp.Location,
		);
		err != nil {
			return nil, err
		}
		chirps = append(chirps, chirp)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chirps, nil
}

func GetChirpByID(id int) (Chirp, error) {
	// a chirp to hold data from the returned row
	var chirp Chirp

	row := db.QueryRow("SELECT * FROM chirps WHERE id = ?", id)
	// Use row.Scan to copy column values into struct fields
	if err := row.Scan(
		&chirp.ID,
		&chirp.Content,
		&chirp.Location,
	);
	err != nil {
		return chirp, err
	}
	return chirp, nil
}