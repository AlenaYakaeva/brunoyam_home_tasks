package tasks

type Task struct {
	TID         string
	Title       string
	Description string
	Status      string // TODO: сделать enum
}
type AddUpdateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required"`
}
