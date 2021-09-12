package mappers

import (
	model "github.com/Senney/todo-app/application/models"
	"github.com/Senney/todo-app/domain/entities"
)

func MapTodoToTodoModel(todo entities.Todo) (model.Todo, error) {
	todoStatus, _ := MapTodoStatusToTodoStatusModel(todo.Status)

	return model.Todo{
		ID:      todo.ID,
		Content: todo.Content,
		Status:  todoStatus,
	}, nil
}

func MapTodoStatusToTodoStatusModel(todoStatus entities.TodoStatus) (model.TodoStatus, error) {
	if todoStatus == entities.TodoStatusCompleted {
		return model.TodoStatusComplete, nil
	}

	if todoStatus == entities.TodoStatusIncomplete {
		return model.TodoStatusIncomplete, nil
	}

	return "Invalid todo status", nil
}
