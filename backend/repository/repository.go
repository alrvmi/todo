package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Task represents a task in the database
type Task struct {
	ID         int
	Title      string
	DueDate    string
	Priority   string
	Completed  bool
	CreatedAt  string
}

// PostgresRepository handles database operations
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository initializes the database
func NewPostgresRepository() *PostgresRepository {
	connStr := "user=postgres dbname=todo sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT,
		due_date TEXT,
		priority TEXT,
		completed BOOLEAN,
		created_at TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}
	return &PostgresRepository{db: db}
}

// SaveTask saves a task to the database
func (r *PostgresRepository) SaveTask(task Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, due_date, priority, completed, created_at) VALUES ($1, $2, $3, $4, $5)",
		task.Title, task.DueDate, task.Priority, task.Completed, task.CreatedAt)
	return err
}

// GetTasks retrieves all tasks
func (r *PostgresRepository) GetTasks() ([]Task, error) {
	rows, err := r.db.Query("SELECT id, title, due_date, priority, completed, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.DueDate, &t.Priority, &t.Completed, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// DeleteTask removes a task by ID
func (r *PostgresRepository) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}

// UpdateTask updates a task's completion status
func (r *PostgresRepository) UpdateTask(id int, completed bool) error {
	_, err := r.db.Exec("UPDATE tasks SET completed = $1 WHERE id = $1", completed, id)
	return err
}