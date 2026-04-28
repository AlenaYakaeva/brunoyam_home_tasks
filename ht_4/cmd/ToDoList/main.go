package main

import (
	"ToDoList/internal/repository/memstorage"
	"ToDoList/internal/server"
	"ToDoList/internal/service/tasks"
	"ToDoList/internal/service/users"
)

func main() {
	repo := memstorage.New()
	usersService := users.New(repo)
	taskService := tasks.New(repo)
	srv := server.New(":8080", usersService, taskService)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
