package app

import (
	"errors"
	"log"
	"net/http"

	"github.com/codescalersinternships/todoapp-omar/models"
	"github.com/gin-gonic/gin"
)

type taskResponse struct {
	Msg  string      `json:"msg"`
	Task models.Task `json:"task"`
}
type multipleTaskResponse struct {
	Msg   string        `json:"msg"`
	Tasks []models.Task `json:"tasks"`
}

type errorResponse struct {
	Err string `json:"err"`
}

// getTasks 		godoc
// @Summary 		Get tasks
// @Description Retrieve a list of tasks from the database.
// @Produce 		application/json
// @Tags 				tasks
// @Success 		200 {object} multipleTaskResponse
// @Failure 		500 {object} errorResponse "internal server error"
// @Router 			/task [get]
func (a *App) getTasks(c *gin.Context) {
	tasks, err := a.DB.GetTasks()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errorResponse{
			Err: "internal server error",
		})
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusOK, multipleTaskResponse{
		Msg:   "tasks are retrieved successfully",
		Tasks: tasks,
	})
}

// AddTask			godoc
// @Summary			Add new task
// @Description	Add a new task to the database.
// @Accept 			application/json
// @Produce 		application/json
// @Param				tags body models.Task true "New task object"
// @Tags				tasks
// @Success			201 {object} taskResponse
// @Failure 		400 {object} errorResponse "failed to read task data"
// @Failure 		500 {object} errorResponse "internal server error"
// @Router			/task [post]
func (a *App) addTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errorResponse{
			Err: "failed to read task data",
		})
		log.Fatal(err)
		return
	}

	newTask, err := a.DB.AddTask(newTask)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errorResponse{
			Err: "internal server error",
		})
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, taskResponse{
		Msg:  "task is created successfully",
		Task: newTask,
	})
}

// editTask 		godoc
// @Summary 		Edit task
// @Description Edit an existing task in the database.
// @Accept 			application/json
// @Produce 		application/json
// @Param 			id path int true "Task ID"
// @Param 			tags body models.Task true "Edited task object"
// @Tags 				tasks
// @Success			200 {object} taskResponse
// @Failure 		400 {object} errorResponse "failed to read task data"
// @Failure 		404 {object} errorResponse "task not found"
// @Failure 		500 {object} errorResponse "internal server error"
// @Router 			/task/{id} [put]
func (a *App) editTask(c *gin.Context) {
	id := c.Param("id")

	var editedTask models.Task
	if err := c.BindJSON(&editedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errorResponse{
			Err: "failed to read task data",
		})
		log.Fatal(err)
		return
	}

	if err := a.DB.EditTask(id, editedTask); err != nil {

		if errors.Is(err, models.ErrTaskNotFound) {
			c.IndentedJSON(http.StatusNotFound, errorResponse{
				Err: "task not found",
			})
			log.Println(err)
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, errorResponse{
			Err: "internal server error",
		})
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusOK, taskResponse{
		Msg:  "task is updated successfully",
		Task: editedTask,
	})
}

// deleteTask 	godoc
// @Summary 		Delete task
// @Description Delete an existing task from the database.
// @Param 			id path int true "Task ID"
// @Tags 				tasks
// @Success 		200
// @Failure			404 {object} errorResponse "task not found"
// @Failure			500 {object} errorResponse "internal server error"
// @Router 			/task/{id} [delete]
func (a *App) deleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := a.DB.DeleteTask(id); err != nil {

		if errors.Is(err, models.ErrTaskNotFound) {
			c.IndentedJSON(http.StatusNotFound, errorResponse{
				Err: "task not found",
			})
			log.Println(err)
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, errorResponse{
			Err: "internal server error",
		})
		log.Fatal(err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"msg": "task is deleted successfully",
	})
}
