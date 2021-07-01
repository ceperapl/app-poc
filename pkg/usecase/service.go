package usecase

//go:generate mockgen -destination=./mocks/mock_service.go -package=mocks github.com/ceperapl/app-poc/pkg/usecase TaskService

import (
	"github.com/ceperapl/app-poc/pkg/models"
	"github.com/ceperapl/app-poc/pkg/repository"
)

// TaskService is a service for managing tasks
type TaskService interface {
	CreateTask(task *models.Task) (*models.Task, error)
	GetTask(id string) (*models.Task, error)
	ListTasks(filterBy string, sortBy string,
		limit int, page int) (tasks []models.Task, count int, err error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTask(cid string) error
}

// NewTaskService returns a task service
func NewTaskService(repos repository.TaskRepository) TaskService {
	return &taskService{repos}
}

type taskService struct {
	repos repository.TaskRepository
}

func (ts *taskService) CreateTask(task *models.Task) (*models.Task, error) {
	if err := ts.repos.CreateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *taskService) GetTask(id string) (*models.Task, error) {
	return ts.repos.GetTask(id)
}

func (ts *taskService) ListTasks(filterBy string, sortBy string,
	limit int, page int) (tasks []models.Task, count int, err error) {
	return ts.repos.ListTasks(filterBy, sortBy, limit, page)
}

func (ts *taskService) UpdateTask(task *models.Task) (*models.Task, error) {
	if err := ts.repos.UpdateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *taskService) DeleteTask(id string) error {
	return ts.repos.DeleteTask(id)
}
