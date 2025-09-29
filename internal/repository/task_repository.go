package repository

import (
	"todo/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db    *gorm.DB
	Task  models.Task
	Tasks []models.Task
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (tr *TaskRepository) GetAll() ([]models.Task, error) {
	r := tr.db.Find(&tr.Tasks)

	return tr.Tasks, r.Error
}

func (tr *TaskRepository) GetByID(id string) (models.Task, error) {
	r := tr.db.First(&tr.Task, id)
	return tr.Task, r.Error
}
