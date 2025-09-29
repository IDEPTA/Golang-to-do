package handlers

import (
	"net/http"
	"todo/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService *service.TaskService
}

func NewTaskHandler(TaskService *service.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: TaskService}
}

func (th *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := th.TaskService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (th *TaskHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	task, err := th.TaskService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (th *TaskHandler) Create(c *gin.Context) {}
