package tasks

import (
	tasksDomain "ToDoList/internal/domain/tasks"
	"ToDoList/internal/service/errors"
	"context"
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	SaveTask(tasksDomain.Task) (string, error)
	GetTasks(string) ([]tasksDomain.Task, error)
	GetTaskByID(string) (tasksDomain.Task, error)
	UpdateTask(tasksDomain.Task, string) (tasksDomain.Task, error)
	DeleteTask(string) error
	DeleteMarkedTasks() error
}
type TaskService struct {
	repo        Repository
	valid       *validator.Validate
	deleteTasks chan struct{}
	batchSize   int
}

func New(repo Repository) *TaskService {
	return &TaskService{
		repo:        repo,
		valid:       validator.New(),
		deleteTasks: make(chan struct{}, 15), //размер канала задан жестко, но можно передавать  размер при создании
		batchSize:   10,
	}
}

func (s *TaskService) AddTask(uid string, req tasksDomain.AddUpdateRequest) (string, error) {

	if err := s.valid.Struct(req); err != nil {
		return "", fmt.Errorf(errors.IncorrectFieldValues, err)
	}

	task := tasksDomain.Task{
		TID:         uuid.NewString(),
		UID:         uid,
		Title:       req.Title,
		Description: req.Description,
		Status:      tasksDomain.ParseStatus(req.Status),
	}

	tid, err := s.repo.SaveTask(task)
	if err != nil {
		return "", err
	}
	return tid, nil
}

func (s *TaskService) GetTasks(uid string) ([]tasksDomain.Task, error) {
	users, err := s.repo.GetTasks(uid)
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

func (s *TaskService) UpdateTask(req tasksDomain.AddUpdateRequest, tid string, uid string) (tasksDomain.Task, error) {

	if err := s.valid.Struct(req); err != nil {
		return tasksDomain.Task{}, fmt.Errorf(errors.IncorrectFieldValues, err)
	}
	task := tasksDomain.Task{
		TID:         tid,
		UID:         uid,
		Title:       req.Title,
		Description: req.Description,
		Status:      tasksDomain.ParseStatus(req.Status),
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
	select {
	case s.deleteTasks <- struct{}{}:
	default:
		log.Info().Msg("Канал заполнен, воркер уже в процессе или скоро начнет очистку")
	}

	return nil

}

func (s *TaskService) StartLazyDeleter(ctx context.Context) {
	count := 0

	for {
		select {
		case <-ctx.Done():
			close(s.deleteTasks)
			for range s.deleteTasks {
				count++
			}
			s.flushAndExecute(&count)
			return

		case <-s.deleteTasks:
			count++
			if count >= s.batchSize {
				s.flushAndExecute(&count)
			}
		}
	}
}

func (s *TaskService) flushAndExecute(count *int) {
	if *count == 0 {
		return
	}

	if err := s.repo.DeleteMarkedTasks(); err != nil {
		log.Error().Err(err).Msg("Ошибка при выполнении DeleteMarkedTasks")
	}
	*count = 0
}
