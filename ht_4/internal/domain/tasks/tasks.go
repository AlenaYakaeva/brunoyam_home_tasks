package tasks

type Task struct {
	ID          string `json:"taskID"`
	Title       string `json:"taskTitle"`
	Description string `json:"taskDescription"`
	Status      string `json:"taskStatus"`
}
