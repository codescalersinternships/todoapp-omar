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

func getTasks(c *gin.Context, client DBClient) {
	tasks, err := client.getTasks()

	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

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
