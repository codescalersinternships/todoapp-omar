package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// NewApp is the factory of App
func NewApp(client DBClient) App {
	return App{client: client}
}

// App is the structure that initializes the entire app
type App struct {
	client DBClient
	Router *gin.Engine
}

// Run runs server
func (a *App) Run(dbFilePath string) error {
	if err := a.client.start(dbFilePath); err != nil {
		return err
	}
	defer a.client.close()

	a.Router = gin.Default()
	a.Router.Use(cors.Default())
	a.setRoutes()
	return a.Router.Run(":8080")
}
