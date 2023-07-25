package internal

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type task struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Is_completed string `json:"is_completed"`
}

func getTasks(c *gin.Context, client *sql.DB) {
	rows, err := client.Query("SELECT * FROM tasks ORDER BY id;")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	data := []task{}

	for rows.Next() {
		t := task{}
		if err := rows.Scan(&t.Id, &t.Title, &t.Is_completed); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		data = append(data, t)
	}
	c.IndentedJSON(http.StatusOK, data)
}

func addTask(c *gin.Context, client *sql.DB) {
	var newTask task
	if err := c.BindJSON(&newTask); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	res, err := client.Exec("INSERT INTO tasks(title) VALUES(?);", newTask.Title)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	newTask.Id = strconv.Itoa(int(id))
	newTask.Is_completed = "false"

	c.IndentedJSON(http.StatusCreated, newTask)
}

func editTask(c *gin.Context, client *sql.DB) {
	id := c.Param("id")

	var editedTask task
	if err := c.BindJSON(&editedTask); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := client.Exec("UPDATE tasks set title = ?, is_completed = ? where id = ?;", editedTask.Title, editedTask.Is_completed, id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, editedTask)
}

func deleteTask(c *gin.Context, client *sql.DB) {
	id := c.Param("id")

	_, err := client.Exec("DELETE FROM tasks where id = ?;", id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
