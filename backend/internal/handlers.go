package internal

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Is_completed bool   `json:"is_completed"`
}

// getTasks 		godoc
// @Summary 		Get tasks
// @Description Retrieve a list of tasks from the database.
// @Produce 		application/json
// @Tags 				tasks
// @Success 		200 {array} task
// @Router 			/task [get]
func getTasks(c *gin.Context, client DBClient) {
	tasks, err := client.getTasks()

	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
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
func addTask(c *gin.Context, client DBClient) {
	var newTask task
	if err := c.BindJSON(&newTask); err != nil {
		log.Fatal(err)
		c.Status(http.StatusBadRequest)
		return
	}

	newTask, err := client.addTask(newTask)

	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusCreated, newTask)
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
func editTask(c *gin.Context, client DBClient) {
	id := c.Param("id")

	var editedTask task
	if err := c.BindJSON(&editedTask); err != nil {
		log.Fatal(err)
		c.Status(http.StatusBadRequest)
		return
	}

	if err := client.editTask(id, editedTask); err != nil {

		if errors.Is(err, errTaskNotFound) {
			c.Status(http.StatusNotFound)
			return
		}

		log.Fatal(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, editedTask)
}

// deleteTask 	godoc
// @Summary 		Delete task
// @Description Delete an existing task from the database.
// @Param 			id path int true "Task ID"
// @Tags 				tasks
// @Success 		200
// @Router 			/task/{id} [delete]
func deleteTask(c *gin.Context, client DBClient) {
	id := c.Param("id")

	if err := client.deleteTask(id); err != nil {

		if errors.Is(err, errTaskNotFound) {
			c.Status(http.StatusNotFound)
			return
		}

		log.Fatal(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
