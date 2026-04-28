package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("Пользователь с таким email уже существует")
	ErrUserNotFound      = errors.New("Пользователь не существует")
)
