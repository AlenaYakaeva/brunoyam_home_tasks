package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

type ToDoList struct {
	Owner    string        `json:"Owner"`
	ListName string        `json:"ListName"`
	Tasks    map[int]*Task `json:"Tasks"`
}

type Task struct {
	Order       int    `json:"taskOrder"`
	IsDone      bool   `json:"taskIsDone"`
	Description string `json:"taskDescription"`
}

var (
	ErrorAddTask    error = errors.New("Задача с таким номером уже существует.")
	ErrorUpdateTask error = errors.New("Невозможно обновить несуществующую задачу.")
	ErrorDeleteTask error = errors.New("Невозможно удалить несуществующую задачу.")
	ErrorSearchTask error = errors.New("Задача с таким номером не существует.")
)

func main() {
	fmt.Println("Start!")
	r := echo.New()

	//r.Use(middleware.Logger())
	tasks := make(map[int]*Task, 10)

	var list ToDoList = ToDoList{
		Owner:    "Иванов Иван",
		ListName: "Уборка",
		Tasks:    tasks,
	}

	//Получить список задач
	r.GET("/tasks", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, list.Tasks)
	})

	//Создать новую задачу в списке
	r.POST("/tasks", func(ctx echo.Context) error {
		var task Task
		err := ctx.Bind(&task)
		if err != nil {
			return ctx.String(http.StatusBadRequest, fmt.Sprint(err))

		}
		_, ok := list.Tasks[task.Order]
		if ok {
			return ctx.String(http.StatusBadRequest, fmt.Sprint(ErrorAddTask))

		}
		list.Tasks[task.Order] = &task

		return ctx.JSON(http.StatusCreated, list)
	})

	//Обновить информацию о задаче по ее ID
	r.PUT("/tasks/:id", func(ctx echo.Context) error {
		Order, _ := strconv.Atoi(ctx.Param("id"))
		var task Task
		err := ctx.Bind(&task)
		if err != nil {
			return ctx.String(http.StatusBadRequest, fmt.Sprint(err))
		}
		_, ok := list.Tasks[Order]
		if !ok {
			return ctx.String(http.StatusBadRequest, fmt.Sprint(ErrorUpdateTask))

		}
		list.Tasks[Order].IsDone = task.IsDone
		list.Tasks[Order].Description = task.Description
		return ctx.JSON(http.StatusAccepted, list)
	})

	//Удалить задачу по ее ID
	r.DELETE("/tasks/:id", func(ctx echo.Context) error {
		Order, _ := strconv.Atoi(ctx.Param("id"))
		_, ok := list.Tasks[Order]
		if !ok {
			return ctx.String(http.StatusBadRequest, fmt.Sprint(ErrorDeleteTask))

		}
		delete(list.Tasks, Order)
		return ctx.JSON(http.StatusAccepted, list)
	})

	//Получить информацию о задаче по ее ID
	r.GET("/tasks/:id", func(ctx echo.Context) error {
		Order, _ := strconv.Atoi(ctx.Param("id"))
		_, ok := list.Tasks[Order]
		if !ok {
			return ctx.String(http.StatusBadRequest, fmt.Sprint(ErrorSearchTask))

		}
		return ctx.JSON(http.StatusAccepted, list.Tasks[Order])
	})

	r.Start(":8080")
}
