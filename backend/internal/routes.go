package internal

import (
	"github.com/gin-gonic/gin"
)

func (a *App) setRoutes() {
	a.Router.GET("task", func(ctx *gin.Context) {
		getTasks(ctx, a.client)
	})
	a.Router.POST("task", func(ctx *gin.Context) {
		addTask(ctx, a.client)
	})
	a.Router.PUT("task/:id", func(ctx *gin.Context) {
		editTask(ctx, a.client)
	})
	a.Router.DELETE("task/:id", func(ctx *gin.Context) {
		deleteTask(ctx, a.client)
	})
}
