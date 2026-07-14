package main

import (
	"ToDoList/internal"
	"ToDoList/internal/repository"
	"ToDoList/internal/server"
	taskService "ToDoList/internal/server/tasks"
	userService "ToDoList/internal/server/users"
	"ToDoList/internal/service/tasks"
	"ToDoList/internal/service/users"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	var wg sync.WaitGroup
	var (
		usersService userService.UserService
		taskService  taskService.TaskService
	)
	//dbDSN := "postgres://postgres:postgresql_pass@localhost:5432/ToDoList?sslmode=disable"
	cfg := internal.ReadConfig()
	cfg.ConfigureLogger()
	repo, err := repository.New(cfg.DBDSN, 5)
	if err != nil {
		usersService = users.New(repo.MemStorage)
		taskService = tasks.New(repo.MemStorage)
	} else {
		usersService = users.New(repo.DB)
		taskService = tasks.New(repo.DB)
		defer repo.DB.Close()
	}

	// Контекст для фоновых воркеров
	workerCtx, cancelWorkers := context.WithCancel(context.Background())

	wg.Add(1)
	// Запускаем фоновый воркер
	go func() {
		defer wg.Done()
		taskService.StartLazyDeleter(workerCtx)
	}()

	srv := server.New(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), usersService, taskService)

	log.Info().Msg("Запускаем сервер")
	go func() {
		if err := srv.Run(); err != nil {
			log.Warn().Msg("Сервер завершил работу")
		}
	}()

	server.WaitForShutdown(srv, 5*time.Second)
	defer log.Info().Msg("Сервис завершил работу")

	cancelWorkers()

	wg.Wait()

}
