package main

import (
	"ToDoList/internal"
	"ToDoList/internal/repository"
	"ToDoList/internal/server"
	taskService "ToDoList/internal/server/tasks"
	userService "ToDoList/internal/server/users"
	"ToDoList/internal/service/tasks"
	"ToDoList/internal/service/users"
	"fmt"
)

func main() {
	var (
		usersService userService.UserService
		taskService  taskService.TaskService
	)
	//dbDSN := "postgres://postgres:postgresql_pass@localhost:5432/ToDoList?sslmode=disable"
	cfg := internal.ReadConfig()
	repo := repository.New(cfg.DBDSN)
	if repo.IsRemote {
		usersService = users.New(repo.DB)
		taskService = tasks.New(repo.DB)
	} else {
		usersService = users.New(repo.MemStorage)
		taskService = tasks.New(repo.MemStorage)
	}

	srv := server.New(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), usersService, taskService)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
