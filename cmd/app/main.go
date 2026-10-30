package main

import (
	"log"
	"todo/internal/container"
	"todo/internal/repositories"
	"todo/internal/routes"
)

func main() {
	db := repositories.DB{}
	db.PostgresConnect()
	log.Println("DB connection established", db.GetDB())

	di := container.NewDI(db.GetDB())

	r := routes.NewRouter(di.TaskHandler,
		di.AuthHandler)
	e := r.SetupRoutes()
	e.Run(":8080")
}
