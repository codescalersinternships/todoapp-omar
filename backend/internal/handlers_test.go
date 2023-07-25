package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	app := NewApp()
	err := app.startDB("./todoapp.db")
	assert.Nil(t, err)
	defer app.closeDB()

	app.Router = gin.Default()
	app.setRoutes()

	t.Run("valid", func(t *testing.T) {
		newTask := map[string]string{
			"title": "task1",
		}
		payload, _ := json.Marshal(newTask)

		req, err := http.NewRequest("POST", "/task/", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		newTask := map[string]int{
			"title": 10,
		}
		payload, _ := json.Marshal(newTask)

		req, err := http.NewRequest("POST", "/task/", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestGetTasks(t *testing.T) {
	app := NewApp()
	err := app.startDB("./todoapp.db")
	assert.Nil(t, err)
	defer app.closeDB()

	app.Router = gin.Default()
	app.setRoutes()

	req, err := http.NewRequest("GET", "/task/", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()

	app.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestEditTask(t *testing.T) {
	app := NewApp()
	err := app.startDB("./todoapp.db")
	assert.Nil(t, err)
	defer app.closeDB()

	app.Router = gin.Default()
	app.setRoutes()

	t.Run("valid", func(t *testing.T) {
		editedTask := map[string]string{
			"id":           "1",
			"title":        "task2",
			"is_completed": "true",
		}
		payload, _ := json.Marshal(editedTask)

		req, err := http.NewRequest("PUT", "/task/1", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		editedTask := map[string]int{
			"id":           1,
			"title":        2,
			"is_completed": 1,
		}
		payload, _ := json.Marshal(editedTask)

		req, err := http.NewRequest("PUT", "/task/1", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestDeleteTask(t *testing.T) {
	app := NewApp()
	err := app.startDB("./todoapp.db")
	assert.Nil(t, err)
	defer app.closeDB()

	app.Router = gin.Default()
	app.setRoutes()

	req, err := http.NewRequest("DELETE", "/task/1", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()

	app.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
