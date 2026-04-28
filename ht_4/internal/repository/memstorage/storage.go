package memstorage

import (
	usersDomain "ToDoList/internal/domain/users"
	errorsRepo "ToDoList/internal/repository/errors"
)

type Storage struct {
	users map[string]usersDomain.User
}

func New() *Storage {
	return &Storage{users: make(map[string]usersDomain.User)}
}

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
