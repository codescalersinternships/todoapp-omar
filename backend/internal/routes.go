package internal

import "github.com/gin-gonic/gin"

// Routes handles the routes of the app
func Routes(route *gin.Engine) {
	task := route.Group("/task")
	{
		task.GET("/", getTasks)
		task.POST("/", addTask)
		task.PUT("/:id", editTask)
		task.DELETE("/:id", deleteTask)
	}
}
