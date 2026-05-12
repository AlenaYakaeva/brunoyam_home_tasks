package main

import (
	"ToDoList/internal/repository/db"
	"ToDoList/internal/server"
	"ToDoList/internal/service/tasks"
	"ToDoList/internal/service/users"
)

func main() {
	dbDSN := "postgres://postgres:postgresql_pass@localhost:5432/ToDoList?sslmode=disable"
	repo, err := db.New(dbDSN) //memstorage.New()
	if err != nil {
		panic(err)
	}
	if err = db.RunMigrations(dbDSN); err != nil {
		panic(err)
	}
	usersService := users.New(repo)
	taskService := tasks.New(repo)
	srv := server.New(":8080", usersService, taskService)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
