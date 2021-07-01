package memory

import (
	"errors"
	"fmt"
	"time"

	"github.com/ceperapl/app-poc/pkg/models"
	"github.com/google/uuid"
)

// TaskRepo implements task repo in memory
type TaskRepo struct {
	tasks []models.Task
}

var (
	errTaskNotFound = errors.New("Task not found")
	errInvalidUUID  = errors.New("Invalid UUID")
)

// NewTaskRepo returns memory repository implementation
func NewTaskRepo() (*TaskRepo, error) {
	return &TaskRepo{tasks: []models.Task{}}, nil
}

// CreateTask creates new task in memory if ID is not exists
func (m *TaskRepo) CreateTask(task *models.Task) error {
	if task.ID == "" {
		task.ID = uuid.New().String()
	} else if _, err := uuid.Parse(task.ID); err != nil {
		return errInvalidUUID
	}
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	for _, t := range m.tasks {
		if t.ID == task.ID {
			return fmt.Errorf("Task ID already exists: %s", task.ID)
		}
	}
	m.tasks = append(m.tasks, *task)
	return nil
}

// GetTask returns task by id
func (m *TaskRepo) GetTask(id string) (*models.Task, error) {
	for _, t := range m.tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errTaskNotFound
}

// ListTasks returns tasks
func (m *TaskRepo) ListTasks(filterBy string, sortBy string, limit int,
	page int) (tasks []models.Task, count int, err error) {
	return m.tasks, len(m.tasks), nil
}

// UpdateTask updates task
func (m *TaskRepo) UpdateTask(task *models.Task) error {
	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = *task
			return nil
		}
	}
	return errTaskNotFound
}

// DeleteTask deletes task
func (m *TaskRepo) DeleteTask(id string) error {
	for i, t := range m.tasks {
		if t.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}
