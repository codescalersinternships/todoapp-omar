package models

import "errors"

// ErrTaskNotFound is task not found error
var ErrTaskNotFound = errors.New("task not found")

// Task model
type Task struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}

// GetTasks get all tasks from db
func (d *DBClient) GetTasks() ([]Task, error) {
	q, err := d.Client.Query("SELECT * FROM tasks ORDER BY id DESC;")
	if err != nil {
		return []Task{}, err
	}
	defer q.Close()

	tasks := []Task{}

	for q.Next() {
		t := Task{}
		if err := q.Scan(&t.ID, &t.Title, &t.IsCompleted); err != nil {
			return []Task{}, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// AddTask adds new task to db
func (d *DBClient) AddTask(newTask Task) (Task, error) {
	response, err := d.Client.Exec("INSERT INTO tasks(title) VALUES(?);", newTask.Title)

	if err != nil {
		return Task{}, err
	}

	id, err := response.LastInsertId()
	if err != nil {
		return Task{}, err
	}

	newTask.ID = uint(id)
	newTask.IsCompleted = false

	return newTask, nil
}

// EditTask edits task
func (d *DBClient) EditTask(id string, editedTask Task) error {
	response, err := d.Client.Exec("UPDATE tasks set title = ?, is_completed = ? where id = ?;", editedTask.Title, editedTask.IsCompleted, id)

	if err != nil {
		return err
	}

	count, err := response.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrTaskNotFound
	}

	return nil
}

// DeleteTask deletes task
func (d *DBClient) DeleteTask(id string) error {
	response, err := d.Client.Exec("DELETE FROM tasks where id = ?;", id)

	if err != nil {
		return err
	}

	count, err := response.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrTaskNotFound
	}

	return nil
}
