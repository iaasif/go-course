package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error
	// DB, err = sql.Open("sqlite3", "api.db")
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	createTables()
}

func createTables() {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create user table: " + err.Error())
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events 
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL, 
		description TEXT NOT NULL, 
		location TEXT NOT NULL, 
		dateTime DATETIME NOT NULL, 
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES user(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table: " + err.Error())
	}
}
