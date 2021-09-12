package todo

import (
	entities "github.com/Senney/todo-app/domain/entities"
)

type TodoProviderPort interface {
	GetTodos() []entities.Todo
}
