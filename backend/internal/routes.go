package internal

import (
	"github.com/gin-gonic/gin"
)

func (a *App) setRoutes() {
	task := a.Router.Group("/task")
	{
		task.GET("/", func(ctx *gin.Context) {
			getTasks(ctx, a.Client)
		})
		task.POST("/", func(ctx *gin.Context) {
			addTask(ctx, a.Client)
		})
		task.PUT("/:id", func(ctx *gin.Context) {
			editTask(ctx, a.Client)
		})
		task.DELETE("/:id", func(ctx *gin.Context) {
			deleteTask(ctx, a.Client)
		})
	}
}
