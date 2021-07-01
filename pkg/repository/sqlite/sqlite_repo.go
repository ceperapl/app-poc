package sqlite

import (
	"github.com/ceperapl/app-poc/pkg/models"
	"github.com/ceperapl/app-poc/pkg/repository/gormmodels"
	"github.com/jinzhu/gorm"
)

type sqliteTaskRepo struct {
	DB *gorm.DB
}

// NewSQLiteTaskRepo returns repository for SQLite
func NewSQLiteTaskRepo(db *gorm.DB) (*sqliteTaskRepo, error) {
	return &sqliteTaskRepo{DB: db}, nil
}

func (p *sqliteTaskRepo) CreateTask(task *models.Task) error {
	gormTask := gormmodels.ToGormTask(task)

	return p.DB.Create(&gormTask).Error
}

func (p *sqliteTaskRepo) GetTask(id string) (*models.Task, error) {
	gormTask := gormmodels.GormTask{}
	if err := p.DB.First(&gormTask, "id = ?", id).Error; err != nil {
		return nil, err
	}
	task := gormTask.ToTask()

	return &task, nil
}

func (p *sqliteTaskRepo) ListTasks(filterBy string, sortBy string, limit int,
	page int) (tasks []*models.Task, count int, err error) {
	return nil, 0, nil
}

func (p *sqliteTaskRepo) UpdateTask(task *models.Task) error {
	return nil
}

func (p *sqliteTaskRepo) DeleteTask(id string) error {
	return nil
}
