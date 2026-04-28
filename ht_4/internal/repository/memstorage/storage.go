package memstorage

import (
	tasksDomain "ToDoList/internal/domain/tasks"
	usersDomain "ToDoList/internal/domain/users"
	errorsRepo "ToDoList/internal/repository/errors"
)

type Storage struct {
	users map[string]usersDomain.User
	tasks map[string]tasksDomain.Task
}

func New() *Storage {
	return &Storage{
		users: make(map[string]usersDomain.User),
		tasks: make(map[string]tasksDomain.Task),
	}
}

// ======Методы работы с пользователем=======
func (s *Storage) SaveUser(user usersDomain.User) error {
	for _, u := range s.users {
		if u.Email == user.Email {
			return errorsRepo.ErrUserAlreadyExists
		}
	}
	s.users[user.UID] = user
	return nil
}

func (s *Storage) GetUsers() ([]usersDomain.User, error) {
	users := make([]usersDomain.User, 0, len(s.users))
	for _, u := range s.users {
		users = append(users, u)
	}
	return users, nil
}

func (s *Storage) GetUserByID(uid string) (usersDomain.User, error) {
	user, err := s.users[uid]
	if !err {
		return usersDomain.User{}, errorsRepo.ErrUserNotFound
	}
	return user, nil
}

func (s *Storage) UpdateUser(user usersDomain.User, uid string) (usersDomain.User, error) {
	_, err := s.users[uid]
	if !err {
		return usersDomain.User{}, errorsRepo.ErrUserNotFound
	}
	s.users[uid] = user
	return user, nil
}

func (s *Storage) DeleteUser(uid string) error {
	_, err := s.users[uid]
	if !err {
		return errorsRepo.ErrUserNotFound
	}
	delete(s.users, uid)
	return nil
}

// ======Методы работы с задачами=======
func (s *Storage) SaveTask(task tasksDomain.Task) error {
	s.tasks[task.TID] = task
	return nil
}

func (s *Storage) GetTasks() ([]tasksDomain.Task, error) {
	tasks := make([]tasksDomain.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (s *Storage) GetTaskByID(tid string) (tasksDomain.Task, error) {
	task, err := s.tasks[tid]
	if !err {
		return tasksDomain.Task{}, errorsRepo.ErrTaskNotFound
	}
	return task, nil
}
func (s *Storage) UpdateTask(task tasksDomain.Task, tid string) (tasksDomain.Task, error) {
	_, err := s.tasks[tid]
	if !err {
		return tasksDomain.Task{}, errorsRepo.ErrTaskNotFound
	}
	s.tasks[tid] = task
	return task, nil
}

func (s *Storage) DeleteTask(tid string) error {
	_, err := s.tasks[tid]
	if !err {
		return errorsRepo.ErrTaskNotFound
	}
	delete(s.tasks, tid)
	return nil
}
