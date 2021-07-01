package repository

//go:generate mockgen -destination=./mocks/mock_repository.go -package=mocks github.com/ceperapl/app-poc/pkg/repository TaskRepository

import (
	"github.com/ceperapl/app-poc/pkg/models"
)

// TaskRepository is an abstract repository layer for data access
type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTask(id string) (*models.Task, error)
	ListTasks(filterBy string, sortBy string, limit int, page int) (tasks []models.Task, count int, err error)
	UpdateTask(task *models.Task) error
	DeleteTask(id string) error
}
