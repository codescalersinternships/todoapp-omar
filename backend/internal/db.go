package internal

import (
	"database/sql"
	_ "embed"
	"errors"
)

var errTaskNotFound = errors.New("task not found")

// DBClient used to start, close and make queries on database
type DBClient struct {
	client *sql.DB
}

//go:embed db/000createTables.sql
var createTableQuery string

func (d *DBClient) start(dbFilePath string) error {
	// connect to sqlite db
	var err error
	if d.client, err = sql.Open("sqlite3", dbFilePath); err != nil {
		return err
	}

	// create tables
	_, err = d.client.Exec(createTableQuery)
	return err
}

func (d *DBClient) close() {
	d.client.Close()
}

func (d *DBClient) getTasks() ([]task, error) {
	rows, err := d.client.Query("SELECT * FROM tasks ORDER BY id DESC;")
	if err != nil {
		return []task{}, err
	}
	defer rows.Close()

	tasks := []task{}

	for rows.Next() {
		t := task{}
		if err := rows.Scan(&t.Id, &t.Title, &t.Is_completed); err != nil {
			return []task{}, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (d *DBClient) addTask(newTask task) (task, error) {
	response, err := d.client.Exec("INSERT INTO tasks(title) VALUES(?);", newTask.Title)

	if err != nil {
		return task{}, err
	}

	id, err := response.LastInsertId()
	if err != nil {
		return task{}, err
	}

	newTask.Id = int(id)
	newTask.Is_completed = false

	return newTask, nil
}

func (d *DBClient) editTask(id string, editedTask task) error {
	response, err := d.client.Exec("UPDATE tasks set title = ?, is_completed = ? where id = ?;", editedTask.Title, editedTask.Is_completed, id)
	if err != nil {
		return err
	}

	count, err := response.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errTaskNotFound
	}

	return nil
}

func (d *DBClient) deleteTask(id string) error {
	response, err := d.client.Exec("DELETE FROM tasks where id = ?;", id)
	if err != nil {
		return err
	}

	count, err := response.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errTaskNotFound
	}

	return nil
}
