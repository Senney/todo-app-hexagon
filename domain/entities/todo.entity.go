package entities

type TodoStatus int

const (
	TodoStatusCompleted TodoStatus = iota
	TodoStatusIncomplete
)

type Todo struct {
	ID      string
	Content string
	Status  TodoStatus
}
