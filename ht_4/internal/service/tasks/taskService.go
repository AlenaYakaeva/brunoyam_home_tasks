package tasks

import (
	tasksDomain "ToDoList/internal/domain/tasks"
	"ToDoList/internal/service/errors"
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	SaveTask(tasksDomain.Task) error
	GetTasks() ([]tasksDomain.Task, error)
	GetTaskByID(string) (tasksDomain.Task, error)
	UpdateTask(tasksDomain.Task, string) (tasksDomain.Task, error)
	DeleteTask(string) error
}
type TaskService struct {
	repo  Repository
	valid *validator.Validate
}

func New(repo Repository) *TaskService {
	return &TaskService{
		repo:  repo,
		valid: validator.New(),
	}
}

func (s *TaskService) AddTask(req tasksDomain.AddUpdateRequest) (string, error) {

	if err := s.valid.Struct(req); err != nil {
		return "", fmt.Errorf(errors.IncorrectFieldValues, err)
	}

	task := tasksDomain.Task{
		TID:         uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	if err := s.repo.SaveTask(task); err != nil {
		return "", err
	}
	return task.TID, nil
}

func (s *TaskService) GetTasks() ([]tasksDomain.Task, error) {
	users, err := s.repo.GetTasks()
	if err != nil {
		return []tasksDomain.Task{}, err
	}
	return users, nil
}

func (s *TaskService) FindTaskByID(tid string) (tasksDomain.Task, error) {
	task, err := s.repo.GetTaskByID(tid)
	if err != nil {
		return tasksDomain.Task{}, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(req tasksDomain.AddUpdateRequest, tid string) (tasksDomain.Task, error) {

	if err := s.valid.Struct(req); err != nil {
		return tasksDomain.Task{}, fmt.Errorf(errors.IncorrectFieldValues, err)
	}
	task := tasksDomain.Task{
		TID:         tid,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	updateTask, err := s.repo.UpdateTask(task, tid)
	if err != nil {
		return tasksDomain.Task{}, err
	}
	return updateTask, nil
}

func (s *TaskService) DeleteTask(tid string) error {
	err := s.repo.DeleteTask(tid)
	if err != nil {
		return err
	}
	return nil
}
