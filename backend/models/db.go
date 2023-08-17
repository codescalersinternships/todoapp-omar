package models

import (
	"database/sql"
	_ "embed"
)

// NewDBClient connects sqlite database and returns DBClient
func NewDBClient(dbFilePath string) (DBClient, error) {
	d := DBClient{}

	// connect database
	var err error
	d.Client, err = sql.Open("sqlite3", dbFilePath)

	return d, err
}

// DBClient used to start, close and make queries on database
type DBClient struct {
	Client *sql.DB
}

//go:embed db/000createTables.sql
var createTableQuery string

// Migrate creates table for the database
func (d *DBClient) Migrate() error {
	_, err := d.Client.Exec(createTableQuery)
	return err
}

// Close closes db connection
func (d *DBClient) Close() {
	d.Client.Close()
}
