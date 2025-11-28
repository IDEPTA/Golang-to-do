package main

import (
	"log"
	"todo/internal/handlers"
	"todo/internal/repositories"
	"todo/internal/routes"
	"todo/internal/services"
)

func main() {
	db := new(repositories.DB)
	db.PostgresConnect()
	log.Println("DB connection established", db.GetDB())
	tr := repositories.NewTaskRepository(db.GetDB())
	ts := services.NewTaskService(tr)
	th := handlers.NewTaskHandler(ts)

	ar := repositories.NewAuthkRepository(db.GetDB())
	as := services.NewAuthService(ar)
	ah := handlers.NewAuthHandler(as)

	r := routes.NewRouter(th, ah)
	e := r.SetupRoutes()
	e.Run(":8080")
}
