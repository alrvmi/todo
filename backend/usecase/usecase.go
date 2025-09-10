package usecase

import (
	"time"
	"todo-app/backend/service"
)

// Task represents a task in the usecase layer
type Task struct {
	ID        int
	Title     string
	DueDate   string
	Priority  string
	Completed bool
	CreatedAt string
}

// TaskUsecase handles application logic
type TaskUsecase struct {
	service *service.TaskService
}

// NewTaskUsecase creates a new usecase instance
func NewTaskUsecase(service *service.TaskService) *TaskUsecase {
	return &TaskUsecase{service: service}
}

// AddTask adds a new task
func (u *TaskUsecase) AddTask(title, dueDate, priority string) error {
	return u.service.AddTask(title, dueDate, priority)
}

// FilterTasks filters tasks by status or date
func (u *TaskUsecase) FilterTasks(status, dateFilter string) ([]Task, error) {
	tasks, err := u.service.GetTasks()
	if err != nil {
		return nil, err
	}
	var filtered []Task
	for _, t := range tasks {
		task := Task{
			ID:        t.ID,
			Title:     t.Title,
			DueDate:   t.DueDate,
			Priority:  t.Priority,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		}
		if status == "all" || (status == "active" && !t.Completed) || (status == "completed" && t.Completed) {
			if dateFilter == "" || matchesDateFilter(t.DueDate, dateFilter) {
				filtered = append(filtered, task)
			}
		}
	}
	return filtered, nil
}

// matchesDateFilter checks if task matches date filter
func matchesDateFilter(dueDate, filter string) bool {
	if dueDate == "" {
		return false
	}
	due, _ := time.Parse(time.RFC3339, dueDate)
	now := time.Now()
	switch filter {
	case "today":
		return due.Year() == now.Year() && due.YearDay() == now.YearDay()
	case "week":
		return due.Before(now.AddDate(0, 0, 7))
	case "overdue":
		return due.Before(now) && !due.IsZero()
	}
	return true
}