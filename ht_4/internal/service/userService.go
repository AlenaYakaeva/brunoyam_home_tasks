package service

import (
	usersDomain "ToDoList/internal/domain/users"
	"ToDoList/internal/service/errors"
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Repository interface {
	SaveUser(usersDomain.User) error
	GetUsers() ([]usersDomain.User, error)
	GetUserByID(string) (usersDomain.User, error)
	UpdateUser(usersDomain.User, string) (usersDomain.User, error)
	DeleteUser(string) error
}
type UserService struct {
	repo  Repository
	valid *validator.Validate
}

func New(repo Repository) *UserService {
	return &UserService{
		repo:  repo,
		valid: validator.New(),
	}
}

func (s *UserService) RegisterUser(req usersDomain.RegisterRequest) (string, error) {

	if err := s.valid.Struct(req); err != nil {
		return "", fmt.Errorf(errors.IncorrectFieldValues, err)
	}

	user := usersDomain.User{
		UID:      uuid.NewString(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.repo.SaveUser(user); err != nil {
		return "", err
	}
	return user.UID, nil
}

func (s *UserService) GetUsers() ([]usersDomain.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return []usersDomain.User{}, err
	}
	return users, nil
}

func (s *UserService) FindUserByID(uid string) (usersDomain.User, error) {
	user, err := s.repo.GetUserByID(uid)
	if err != nil {
		return usersDomain.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(req usersDomain.UpdateRequest, uid string) (usersDomain.User, error) {

	if err := s.valid.Struct(req); err != nil {
		return usersDomain.User{}, fmt.Errorf(errors.IncorrectFieldValues, err)
	}
	user, err := s.repo.GetUserByID(uid)
	if err != nil {
		return usersDomain.User{}, err
	}
	user.Name = req.Name
	user.Password = req.Password

	updateUser, err := s.repo.UpdateUser(user, uid)
	if err != nil {
		return usersDomain.User{}, err
	}
	return updateUser, nil
}

func (s *UserService) DeleteUser(uid string) error {
	err := s.repo.DeleteUser(uid)
	if err != nil {
		return err
	}
	return nil
}
