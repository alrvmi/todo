package service

import (
	"errors"
	"time"
	"todo-app/backend/repository"
)

// TaskService handles business logic
type TaskService struct {
	repo *repository.PostgresRepository
}

// NewTaskService creates a new service instance
func NewTaskService(repo *repository.PostgresRepository) *TaskService {
	return &TaskService{repo: repo}
}

// AddTask adds a new task
func (s *TaskService) AddTask(title, dueDate, priority string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	task := repository.Task{
		Title:     title,
		DueDate:   dueDate,
		Priority:  priority,
		Completed: false,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	return s.repo.SaveTask(task)
}

// GetTasks retrieves all tasks
func (s *TaskService) GetTasks() ([]repository.Task, error) {
	return s.repo.GetTasks()
}