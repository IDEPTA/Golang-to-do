package routes

import (
	"todo/internal/handlers"
	"todo/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	taskHandler *handlers.TaskHandler
	authHandler *handlers.AuthHandler
}

func NewRouter(taskHandler *handlers.TaskHandler, authHandler *handlers.AuthHandler) *Router {
	return &Router{
		taskHandler: taskHandler,
		authHandler: authHandler,
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	e := gin.Default()

	api := e.Group("/api")

	tasks := api.Group("/tasks")
	tasks.Use(middleware.AuthMiddleware(r.authHandler.As.Ar))
	{
		tasks.GET("/", r.taskHandler.GetAll)
		tasks.GET("/:id", r.taskHandler.GetByID)
		tasks.PUT("/:id", r.taskHandler.Update)
		tasks.POST("/", r.taskHandler.Create)
		tasks.DELETE("/:id", r.taskHandler.Delete)
	}

	auth := api.Group("/auth")
	{
		auth.POST("/login", r.authHandler.Login)
		auth.POST("/register", r.authHandler.Register)
		auth.GET("/me", middleware.AuthMiddleware(r.authHandler.As.Ar), r.authHandler.Me)
	}
	return e
}
