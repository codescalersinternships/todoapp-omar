package app

import (
	"github.com/codescalersinternships/todoapp-omar/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewApp is the factory of App
func NewApp(dbFilePath string) (App, error) {
	db, err := models.NewDBClient(dbFilePath)
	if err != nil {
		return App{}, err
	}

	if err := db.Migrate(); err != nil {
		return App{}, err
	}

	return App{DB: db, Router: gin.Default()}, nil
}

// App is the structure that initializes the entire app
type App struct {
	DB     models.DBClient
	Router *gin.Engine
}

// Run runs server
func (a *App) Run(dbFilePath string) error {
	defer a.DB.Close()

	a.Router.Use(cors.Default())
	a.registerRoutes()

	return a.Router.Run(":8080")
}

func (a *App) registerRoutes() {
	a.Router.GET("task", a.getTasks)
	a.Router.POST("task", a.addTask)
	a.Router.PUT("task/:id", a.editTask)
	a.Router.DELETE("task/:id", a.deleteTask)
	// swagger docs endpoint
	a.Router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
