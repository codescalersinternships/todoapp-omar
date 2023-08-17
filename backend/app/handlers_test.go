package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/codescalersinternships/todoapp-omar/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	dir := t.TempDir()
	app, err := NewApp(filepath.Join(dir, "test.db"))
	assert.Nil(t, err)
	defer app.DB.Close()

	app.Router = gin.Default()
	app.registerRoutes()

	newTask := map[string]string{
		"title": "task1",
	}
	payload, err := json.Marshal(newTask)
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(payload))
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	app.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	// test changes in db
	if w.Code == http.StatusCreated {
		ts, err := app.DB.GetTasks()
		assert.Nil(t, err)
		assert.Equal(t, uint(1), ts[0].ID)
		assert.Equal(t, "task1", ts[0].Title)
		assert.Equal(t, false, ts[0].IsCompleted)
	}
}

func TestGetTasks(t *testing.T) {
	dir := t.TempDir()
	app, err := NewApp(filepath.Join(dir, "test.db"))
	assert.Nil(t, err)
	defer app.DB.Close()

	app.Router = gin.Default()
	app.registerRoutes()

	// pre
	tsk := models.Task{Title: "task1"}
	_, err = app.DB.AddTask(tsk)
	assert.Nil(t, err)

	// test
	req, err := http.NewRequest("GET", "/task", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()

	app.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// test response body with db
	if w.Code == http.StatusOK {
		got := strings.ReplaceAll(w.Body.String(), " ", "")
		got = strings.ReplaceAll(got, "\n", "")

		resp, err := app.DB.GetTasks()
		assert.Nil(t, err)
		EncodedWant, err := json.Marshal(gin.H{
			"msg":   "tasks are retrieved successfully",
			"tasks": resp,
		})
		assert.Nil(t, err)
		want := strings.ReplaceAll(string(EncodedWant), " ", "")

		assert.Equal(t, string(want), got)
	}
}

func TestEditTask(t *testing.T) {
	dir := t.TempDir()
	app, err := NewApp(filepath.Join(dir, "test.db"))
	assert.Nil(t, err)
	defer app.DB.Close()

	app.Router = gin.Default()
	app.registerRoutes()

	t.Run("valid", func(t *testing.T) {
		defer func() {
			_, err := app.DB.Client.Exec("DELETE FROM tasks")
			assert.Nil(t, err)
		}()

		// pre
		tsk := models.Task{Title: "task1"}
		_, err = app.DB.AddTask(tsk)
		assert.Nil(t, err)

		editedTask := models.Task{ID: 1, Title: "task2", IsCompleted: true}
		payload, err := json.Marshal(editedTask)
		assert.Nil(t, err)

		req, err := http.NewRequest("PUT", "/task/1", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		// test changes in db
		if w.Code == http.StatusCreated {
			ts, err := app.DB.GetTasks()
			assert.Nil(t, err)
			assert.Equal(t, uint(1), ts[0].ID)
			assert.Equal(t, "task2", ts[0].Title)
			assert.Equal(t, true, ts[0].IsCompleted)
		}
	})

	t.Run("not found", func(t *testing.T) {
		defer func() {
			_, err := app.DB.Client.Exec("DELETE FROM tasks")
			assert.Nil(t, err)
		}()

		editedTask := models.Task{ID: 2, Title: "task2", IsCompleted: true}
		payload, err := json.Marshal(editedTask)
		assert.Nil(t, err)

		req, err := http.NewRequest("PUT", "/task/1", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("invalid format", func(t *testing.T) {
		defer func() {
			_, err := app.DB.Client.Exec("DELETE FROM tasks")
			assert.Nil(t, err)
		}()

		// pre
		tsk := models.Task{Title: "task1"}
		_, err = app.DB.AddTask(tsk)
		assert.Nil(t, err)

		editedTask := models.Task{ID: 2, IsCompleted: true}
		payload, err := json.Marshal(editedTask)
		assert.Nil(t, err)

		req, err := http.NewRequest("PUT", "/task/1", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteTask(t *testing.T) {
	dir := t.TempDir()
	app, err := NewApp(filepath.Join(dir, "test.db"))
	assert.Nil(t, err)
	defer app.DB.Close()

	app.Router = gin.Default()
	app.registerRoutes()

	t.Run("valid", func(t *testing.T) {
		defer func() {
			_, err := app.DB.Client.Exec("DELETE FROM tasks")
			assert.Nil(t, err)
		}()

		// pre
		tsk := models.Task{Title: "task1"}
		_, err = app.DB.AddTask(tsk)
		assert.Nil(t, err)

		req, err := http.NewRequest("DELETE", "/task/1", nil)
		assert.Nil(t, err)

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("not found", func(t *testing.T) {
		defer func() {
			_, err := app.DB.Client.Exec("DELETE FROM tasks")
			assert.Nil(t, err)
		}()

		req, err := http.NewRequest("DELETE", "/task/1", nil)
		assert.Nil(t, err)

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
