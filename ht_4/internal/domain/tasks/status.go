package tasks

type Status int

const (
	New = iota
	InProgress
	Done
)

var statuses = []string{"Новая", "В процессе", "Завершена"}

func (s Status) String() string {
	return statuses[s]
}
func ParseStatus(status string) Status {
	switch status {
	case "В процессе":
		return InProgress
	case "Завершена":
		return Done
	default:
		return New
	}
}
