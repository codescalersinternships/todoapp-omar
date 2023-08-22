package app

import (
	"errors"
	"fmt"

	"github.com/codescalersinternships/todoapp-omar/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ErrInvalidPort is invalid port number error
var ErrInvalidPort = errors.New("invalid port number, insert a port number from 1 to 65535")

// NewApp is the factory of App
func NewApp(dbFilePath string, port int) (App, error) {
	if port < 1 || port > 65535 {
		return App{}, ErrInvalidPort
	}

	return App{
		DB:     models.NewDBClient(dbFilePath),
		Router: gin.Default(),
		Port:   port,
	}, nil
}

// App is the structure that initializes the entire app
type App struct {
	DB     models.DBClient
	Router *gin.Engine
	Port   int
}

// Run runs server
func (a *App) Run() error {
	if err := a.DB.Connect(); err != nil {
		return err
	}
	defer a.DB.Close()

	if err := a.DB.Migrate(); err != nil {
		return err
	}

	a.registerRoutes()

	return a.Router.Run(fmt.Sprintf(":%d", a.Port))
}

func (a *App) registerRoutes() {
	a.Router.Use(cors.Default())

	a.Router.GET("task", a.getTasks)
	a.Router.POST("task", a.addTask)
	a.Router.PUT("task/:id", a.editTask)
	a.Router.DELETE("task/:id", a.deleteTask)
	// swagger docs endpoint
	a.Router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
