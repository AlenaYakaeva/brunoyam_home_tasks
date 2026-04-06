package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	r := fiber.New()

	tasks := make(map[int]*Task, 10)

	var list ToDoList = ToDoList{
		Owner:    "Иванов Иван",
		ListName: "Уборка",
		Tasks:    tasks,
	}

	//Получить список задач
	r.Get("/tasks", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(list.Tasks)
	})

	//Создать новую задачу в списке
	r.Post("/tasks", func(ctx *fiber.Ctx) error {
		var task Task
		err := ctx.BodyParser(&task)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})

		}
		_, ok := list.Tasks[task.Order]
		if ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrorAddTask})

		}
		list.Tasks[task.Order] = &task

		return ctx.Status(fiber.StatusCreated).JSON(list)
	})

	//Обновить информацию о задаче по ее ID
	r.Put("/tasks/:id", func(ctx *fiber.Ctx) error {
		Order, _ := strconv.Atoi(ctx.Params("id"))
		var task Task
		err := ctx.BodyParser(&task)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}
		_, ok := list.Tasks[Order]
		if !ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrorUpdateTask})
		}
		list.Tasks[Order].IsDone = task.IsDone
		list.Tasks[Order].Description = task.Description
		return ctx.Status(fiber.StatusCreated).JSON(list)

	})
	//Удалить задачу по ее ID
	r.Delete("/tasks/:id", func(ctx *fiber.Ctx) error {
		Order, _ := strconv.Atoi(ctx.Params("id"))

		_, ok := list.Tasks[Order]
		if !ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error2": ErrorDeleteTask})

		}
		delete(list.Tasks, Order)
		return ctx.Status(fiber.StatusAccepted).JSON(list)

	})
	//Получить информацию о задаче по ее ID
	r.Get("/tasks/:id", func(ctx *fiber.Ctx) error {
		Order, _ := strconv.Atoi(ctx.Params("id"))
		_, ok := list.Tasks[Order]
		if !ok {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrorSearchTask})

		}
		return ctx.Status(fiber.StatusAccepted).JSON(list.Tasks[Order])
	})

	r.Listen(":8080")
}
