package errors

import "errors"

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrInvalidJWT         = errors.New("Invalid JWT")
)

const (
	IncorrectFieldValues = "Incorrect field values: %w"
)
