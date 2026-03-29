package db

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	sql.Open("sqlite3", "api.db")
}
