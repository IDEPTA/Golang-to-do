package container

import (
	"todo/internal/handlers"
	"todo/internal/repositories"
	"todo/internal/services"

	"gorm.io/gorm"
)

type DI struct {
	TaskHandler *handlers.TaskHandler
	AuthHandler *handlers.AuthHandler
}

func NewDI(db *gorm.DB) *DI {
	//Зависимости для группы Task
	tr := repositories.NewTaskRepository(db)
	ts := services.NewTaskService(tr)
	th := handlers.NewTaskHandler(ts)

	//Зависимости для группы Auth
	ar := repositories.NewAuthRepository(db)
	as := services.NewAuthService(ar)
	ah := handlers.NewAuthHandler(as)

	return &DI{
		TaskHandler: th,
		AuthHandler: ah,
	}
}
