package repository

import (
	"todo/internal/models"
	"todo/internal/requests"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db    *gorm.DB
	Tasks []models.Task
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (tr *TaskRepository) GetAll() ([]models.Task, error) {
	r := tr.db.Find(&tr.Tasks)

	return tr.Tasks, r.Error
}

func (tr *TaskRepository) GetByID(id int) (models.Task, error) {
	var task models.Task
	r := tr.db.First(&task, id)
	return task, r.Error
}

func (tr *TaskRepository) Create(task models.Task) (models.Task, error) {
	r := tr.db.Create(&task)

	return task, r.Error
}

func (tr *TaskRepository) Update(id int, task requests.TaskRequest) (models.Task, error) {
	ut, err := tr.GetByID(id)
	if err != nil {
		return models.Task{}, err
	}

	ut.Title = task.Title
	ut.Description = task.Description
	ut.Completed = task.Completed
	ut.DateTimeTask = task.DateTimeTask
	ut.UserID = task.UserID
	r := tr.db.Save(&ut)

	return ut, r.Error
}

func (tr *TaskRepository) Delete(id int) error {
	dt, err := tr.GetByID(id)
	if err != nil {
		return err
	}

	r := tr.db.Delete(&dt, id)

	return r.Error
}
