package tasks

import (
	tasksDomain "ToDoList/internal/domain/tasks"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskService interface {
	AddTask(tasksDomain.AddUpdateRequest) (string, error)
	GetTasks() ([]tasksDomain.Task, error)
	FindTaskByID(string) (tasksDomain.Task, error)
	UpdateTask(tasksDomain.AddUpdateRequest, string) (tasksDomain.Task, error)
	DeleteTask(string) error
}

type TaskHandler struct {
	taskService TaskService
}

func NewTaskHandler(taskService TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (uh *TaskHandler) AddTask(ctx *gin.Context) {
	var req tasksDomain.AddUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tid, err := uh.taskService.AddTask(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tid": tid})
}

func (uh *TaskHandler) GetTasks(ctx *gin.Context) {
	tasks, err := uh.taskService.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (uh *TaskHandler) FindTaskByID(ctx *gin.Context) {
	tid := ctx.Param("id")
	task, err := uh.taskService.FindTaskByID(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (uh *TaskHandler) UpdateTask(ctx *gin.Context) {
	tid := ctx.Param("id")
	var req tasksDomain.AddUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := uh.taskService.UpdateTask(req, tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}
func (uh *TaskHandler) DeleteTask(ctx *gin.Context) {
	uid := ctx.Param("id")
	err := uh.taskService.DeleteTask(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
