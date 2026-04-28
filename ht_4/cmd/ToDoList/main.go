package main

import (
	"ToDoList/internal/repository/memstorage"
	"ToDoList/internal/server"
	"ToDoList/internal/service"
)

func main() {
	repo := memstorage.New()
	usersService := service.New(repo)
	srv := server.New(":8080", usersService)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
