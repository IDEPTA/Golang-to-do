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
	tr := repository.NewTaskRepository(db.GetDB())
	ts := service.NewTaskService(tr)
	th := handlers.NewTaskHandler(ts)

	ar := repository.NewAuthkRepository(db.GetDB())
	as := service.NewAuthService(ar)
	ah := handlers.NewAuthHandler(as)

	r := routes.NewRouter(th, ah)
	e := r.SetupRoutes()
	e.Run(":8080")
}
