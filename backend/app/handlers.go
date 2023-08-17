package app

import (
	"errors"
	"log"
	"net/http"

	"github.com/codescalersinternships/todoapp-omar/models"
	"github.com/gin-gonic/gin"
)

// getTasks 		godoc
// @Summary 		Get tasks
// @Description Retrieve a list of tasks from the database.
// @Produce 		application/json
// @Tags 				tasks
// @Success 		200 {array} task
// @Router 			/task [get]
func (a *App) getTasks(c *gin.Context) {
	tasks, err := a.DB.GetTasks()

	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"err": "internal server error",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"msg":   "tasks are retrieved successfully",
		"tasks": tasks,
	})
}

// AddTask			godoc
// @Summary			Add new task
// @Description	Add a new task to the database.
// @Accept 			application/json
// @Produce 		application/json
// @Param				tags body task true "New task object"
// @Tags				tasks
// @Success			201 {object} task
// @Router			/task [post]
func (a *App) addTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": "failed to read task data",
		})
		return
	}

	newTask, err := a.DB.AddTask(newTask)

	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"err": "internal server error",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"msg":  "task is created successfully",
		"task": newTask,
	})
}

// editTask 		godoc
// @Summary 		Edit task
// @Description Edit an existing task in the database.
// @Accept 			application/json
// @Produce 		application/json
// @Param 			id path int true "Task ID"
// @Param 			tags body task true "Edited task object"
// @Tags 				tasks
// @Success 		200 {object} task
// @Router 			/task/{id} [put]
func (a *App) editTask(c *gin.Context) {
	id := c.Param("id")

	var editedTask models.Task
	if err := c.BindJSON(&editedTask); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": "failed to read task data",
		})
		return
	}

	if err := a.DB.EditTask(id, editedTask); err != nil {

		if errors.Is(err, models.ErrTaskNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"err": "task not found",
			})
			return
		}

		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"err": "internal server error",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"msg":  "task is updated successfully",
		"task": editedTask,
	})
}

// deleteTask 	godoc
// @Summary 		Delete task
// @Description Delete an existing task from the database.
// @Param 			id path int true "Task ID"
// @Tags 				tasks
// @Success 		200
// @Router 			/task/{id} [delete]
func (a *App) deleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := a.DB.DeleteTask(id); err != nil {

		if errors.Is(err, models.ErrTaskNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"err": "task not found",
			})
			return
		}

		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"err": "internal server error",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"msg": "task is deleted successfully",
	})
}
