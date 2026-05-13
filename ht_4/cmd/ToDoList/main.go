package main

import (
	"ToDoList/internal/repository"
	"ToDoList/internal/server"
	taskService "ToDoList/internal/server/tasks"
	userService "ToDoList/internal/server/users"
	"ToDoList/internal/service/tasks"
	"ToDoList/internal/service/users"
)

func main() {
	var (
		usersService userService.UserService
		taskService  taskService.TaskService
	)
	dbDSN := "postgres://postgres:postgresql_pass@localhost:5432/ToDoList?sslmode=disable"
	repo := repository.New(dbDSN)
	if repo.IsRemote {
		usersService = users.New(repo.DB)
		taskService = tasks.New(repo.DB)
	} else {
		usersService = users.New(repo.MemStorage)
		taskService = tasks.New(repo.MemStorage)
	}

	srv := server.New(":8080", usersService, taskService)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
