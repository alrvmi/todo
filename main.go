package main

import (
	"context"
	"todo-app/backend/usecase"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	usecase *usecase.TaskUsecase
}

// NewApp creates a new App application struct
func NewApp() *App {
	repo := repository.NewPostgresRepository() // Initialize PostgreSQL repo
	service := service.NewTaskService(repo)
	usecase := usecase.NewTaskUsecase(service)
	return &App{usecase: usecase}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// AddTask adds a new task
func (a *App) AddTask(title string, dueDate string, priority string) error {
	return a.usecase.AddTask(title, dueDate, priority)
}

// GetTasks retrieves all tasks
func (a *App) GetTasks() ([]usecase.Task, error) {
	return a.usecase.GetTasks()
}

// DeleteTask removes a task
func (a *App) DeleteTask(id int) error {
	return a.usecase.DeleteTask(id)
}

// ToggleTaskCompletion toggles task completion status
func (a *App) ToggleTaskCompletion(id int) error {
	return a.usecase.ToggleTaskCompletion(id)
}

// FilterTasks filters tasks by status or date
func (a *App) FilterTasks(status string, dateFilter string) ([]usecase.Task, error) {
	return a.usecase.FilterTasks(status, dateFilter)
}

// SortTasks sorts tasks by date or priority
func (a *App) SortTasks(sortBy string) ([]usecase.Task, error) {
	return a.usecase.SortTasks(sortBy)
}

func main() {
	app := NewApp()
	err := wails.Run(&wails.Options{
		Title:     "To-Do List",
		Width:     800,
		Height:    600,
		AssetDir:  "./frontend/dist",
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}