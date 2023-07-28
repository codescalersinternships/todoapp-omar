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
	app := NewApp(DBClient{})
	err := app.client.start("./todoapp.db")
	assert.Nil(t, err)
	defer app.client.close()

	app.Router = gin.Default()
	app.setRoutes()

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

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}

func TestGetTasks(t *testing.T) {
	app := NewApp(DBClient{})
	err := app.client.start("./todoapp.db")
	assert.Nil(t, err)
	defer app.client.close()

	app.Router = gin.Default()
	app.setRoutes()

	req, err := http.NewRequest("GET", "/task", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()

	app.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestEditTask(t *testing.T) {
	app := NewApp(DBClient{})
	err := app.client.start("./todoapp.db")
	assert.Nil(t, err)
	defer app.client.close()

	app.Router = gin.Default()
	app.setRoutes()

	t.Run("valid", func(t *testing.T) {
		editedTask := task{Id: 1, Title: "task2", Is_completed: true}
		payload, err := json.Marshal(editedTask)
		assert.Nil(t, err)

		req, err := http.NewRequest("PUT", "/task/1", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("not found", func(t *testing.T) {
		editedTask := task{Id: 2, Title: "task2", Is_completed: true}
		payload, err := json.Marshal(editedTask)
		assert.Nil(t, err)

		req, err := http.NewRequest("PUT", "/task/-1", bytes.NewBuffer(payload))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestDeleteTask(t *testing.T) {
	app := NewApp(DBClient{})
	err := app.client.start("./todoapp.db")
	assert.Nil(t, err)
	defer app.client.close()

	app.Router = gin.Default()
	app.setRoutes()

	t.Run("valid", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/task/1", nil)
		assert.Nil(t, err)

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("not found", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/task/-1", nil)
		assert.Nil(t, err)

		w := httptest.NewRecorder()

		app.Router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, w.Code)
		}
	})
}
