package todo

import entities "github.com/Senney/todo-app/domain/entities"

type TodoProviderAdapter struct {
	Todos []entities.Todo
}

func (adapter *TodoProviderAdapter) GetTodos() []entities.Todo {
	return adapter.Todos
}
