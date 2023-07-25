package internal

import (
	"database/sql"

	_ "embed"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// NewApp is the factory of App
func NewApp() App {
	return App{}
}

// App is the structure that initializes the entire app
type App struct {
	Client *sql.DB
	Router *gin.Engine
}

//go:embed db/000createTables.sql
var createTableQuery string

func (a *App) startDB(dbFilePath string) error {
	// connect to sqlite db
	var err error
	if a.Client, err = sql.Open("sqlite3", dbFilePath); err != nil {
		return err
	}

	// create tables
	_, err = a.Client.Exec(createTableQuery)
	return err
}

func (a *App) closeDB() {
	a.Client.Close()
}

// Run runs server
func (a *App) Run(dbFilePath string) error {
	if err := a.startDB(dbFilePath); err != nil {
		return err
	}
	defer a.closeDB()

	a.Router = gin.Default()
	a.setRoutes()
	return a.Router.Run(":8080")
}
