package models

import (
	"database/sql"
	_ "embed"
)

// NewDBClient is the factory of DBClient
func NewDBClient(dbFilePath string) DBClient {
	return DBClient{FilePath: dbFilePath}
}

// DBClient used to start, close and make queries on database
type DBClient struct {
	Client   *sql.DB
	FilePath string
}

// Connect connects sqlite database
func (d *DBClient) Connect() error {
	var err error
	d.Client, err = sql.Open("sqlite3", d.FilePath)
	return err
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
