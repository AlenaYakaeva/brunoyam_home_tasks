package tasks

type Task struct {
	TID         string
	UID         string
	Title       string
	Description string
	Status      Status
}
type AddUpdateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
