package internal

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Client is a client of sqlite database
var Client *sql.DB

// SetupDB connects to database and create needed tables
func SetupDB() error {
	// connect to sqlite db
	var err error
	if Client, err = sql.Open("sqlite3", "todoapp.db"); err != nil {
		return err
	}

	return createTable()
}

func createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		is_completed INTEGER DEFAULT 0
	)
	`

	_, err := Client.Exec(query)
	return err
}
