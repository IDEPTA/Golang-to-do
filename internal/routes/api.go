package routes

import (
	"todo/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	taskHandler *handlers.TaskHandler
}

func NewRouter(taskHandler *handlers.TaskHandler) *Router {
	return &Router{
		taskHandler: taskHandler}
}

func (r *Router) SetupRoutes() *gin.Engine {
	e := gin.Default()

	api := e.Group("/api")
	{
		api.GET("/tasks", r.taskHandler.GetAll)
		api.GET("/tasks/:id", r.taskHandler.GetByID)
		api.PUT("/tasks/:id", r.taskHandler.Update)
		api.POST("/tasks", r.taskHandler.Create)
		api.DELETE("/tasks/:id", r.taskHandler.Delete)
	}

	return e
}
