package server

import (
	"ToDoList/internal/server/tasks"
	"ToDoList/internal/server/users"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server
}

func New(addr string, usersService users.UserService, taskService tasks.TaskService) *Server {
	srv := &http.Server{
		Addr: addr,
	}
	uh := users.NewUserHandler(usersService)
	th := tasks.NewTaskHandler(taskService)
	r := configureRouter(uh, th)
	srv.Handler = r

	return &Server{
		srv: srv,
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func configureRouter(uh *users.UserHandler, th *tasks.TaskHandler) *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	users.POST("/", uh.Register)
	users.GET("/", uh.GetUsers)
	users.GET("/:id", uh.FindUserByID)
	users.PUT("/:id", uh.UpdateUser)
	users.DELETE("/:id", uh.DeleteUser)

	tasks := r.Group("/tasks")
	tasks.POST("/", th.AddTask)
	tasks.GET("/", th.GetTasks)
	tasks.GET("/:id", th.FindTaskByID)
	tasks.PUT("/:id", th.UpdateTask)
	tasks.DELETE("/:id", th.DeleteTask)
	return r
}
