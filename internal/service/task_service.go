package service

import (
	"todo/internal/models"
	"todo/internal/repository"
)

type TaskService struct {
	TaskRepository *repository.TaskRepository
}

func NewTaskService(TaskRepository *repository.TaskRepository) *TaskService {
	return &TaskService{TaskRepository: TaskRepository}
}

func (ts *TaskService) GetAll() ([]models.Task, error) {
	t, err := ts.TaskRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (th *TaskService) GetByID(id string) (models.Task, error) {
	task, err := th.TaskRepository.GetByID(id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
