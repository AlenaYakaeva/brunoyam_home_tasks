package db

import (
	"ToDoList/internal/domain/users"
	"ToDoList/internal/repository/errors"
	"context"
	"time"
)

func (s *Storage) GetUsers() ([]users.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.conn.Query(ctx, "SELECT uid,	name, email, password FROM users where deleted=false")
	if err != nil {
		return nil, err
	}

	var userList []users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.UID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}

	return userList, nil
}

func (s *Storage) GetUserByID(uid string) (users.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user users.User
	err := s.conn.QueryRow(ctx, "SELECT uid, name, email, password FROM users WHERE uid = $1 and deleted=false", uid).Scan(&user.UID, &user.Name, &user.Email, &user.Password)
	//TODO ошибка не найденного пользователя
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}

func (s *Storage) GetUserByEmail(email string) (users.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user users.User
	err := s.conn.QueryRow(ctx, "SELECT uid, name, email, password FROM users WHERE email = $1 and deleted=false", email).Scan(&user.UID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}

func (s *Storage) SaveUser(user users.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.GetUserByEmail(user.Email)
	if err == nil {
		return "", errors.ErrUserAlreadyExists
	}

	var uid string
	err = s.conn.QueryRow(ctx,
		`INSERT INTO users (
		uid, 
		name, 
		email, 
		password) 
		VALUES ($1, $2, $3, $4) RETURNING uid`,
		user.UID,
		user.Name,
		user.Email,
		user.Password).Scan(&uid)
	if err != nil {
		return "", err
	}
	return uid, nil
}

func (s *Storage) UpdateUser(user users.User, uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.GetUserByID(uid)
	if err != nil {
		return errors.ErrUserNotFound
	}

	_, err = s.conn.Exec(ctx, "UPDATE users SET name=$1, password=$2 WHERE uid=$3 and deleted=false", user.Name, user.Password, uid)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteUser(uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.GetUserByID(uid)
	if err != nil {
		return errors.ErrUserNotFound
	}

	_, err = s.conn.Exec(ctx, "UPDATE users SET deleted=true WHERE uid=$1", uid)
	if err != nil {
		return err
	}
	return nil
}
