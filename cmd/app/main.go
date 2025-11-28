package main

import (
	"log"
	"todo/internal/handlers"
	"todo/internal/repository"
	"todo/internal/routes"
	"todo/internal/service"
)

func main() {
	db := new(repository.DB)
	db.PostgresConnect()
	log.Println("DB connection established", db.GetDB())
	// Потом убрать и использовать инъекцию зависимостей
	tr := repository.NewTaskRepository(db.GetDB())
	ts := service.NewTaskService(tr)
	th := handlers.NewTaskHandler(ts)

	r := routes.NewRouter(th)
	e := r.SetupRoutes()
	e.Run(":8080")
}
