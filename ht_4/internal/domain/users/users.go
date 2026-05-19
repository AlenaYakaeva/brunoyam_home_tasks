package users

type User struct {
	UID      string
	Name     string
	Email    string
	Password string
}

type RegisterRequest struct {
	Name     string `json: "name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json "password" validate:"required, min=8"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json "password" validate:"required, min=8"`
}

type UpdateRequest struct {
	Name     string `json: "name" validate:"required"`
	Password string `json "password" validate:"required, min=8"`
}
