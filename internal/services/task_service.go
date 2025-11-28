package services

import (
	"todo/internal/models"
	"todo/internal/repositories"
	"todo/internal/requests"
)

type TaskService struct {
	TaskRepository *repositories.TaskRepository
}

func NewTaskService(TaskRepository *repositories.TaskRepository) *TaskService {
	return &TaskService{TaskRepository: TaskRepository}
}

func (ts *TaskService) GetAll() ([]models.Task, error) {
	t, err := ts.TaskRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (ts *TaskService) GetByID(id int) (models.Task, error) {
	task, err := ts.TaskRepository.GetByID(id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (ts *TaskService) Create(task requests.TaskRequest) (models.Task, error) {
	nt := models.Task{
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		DateTimeTask: task.DateTimeTask,
		UserID:       task.UserID,
	}

	createdTask, err := ts.TaskRepository.Create(nt)
	if err != nil {

		return models.Task{}, err
	}

	return createdTask, nil
}

func (ts *TaskService) Update(id int, task requests.TaskRequest) (models.Task, error) {
	ut, err := ts.TaskRepository.Update(id, task)
	if err != nil {
		return models.Task{}, err
	}

	return ut, nil
}

func (ts *TaskService) Delete(id int) error {
	err := ts.TaskRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
