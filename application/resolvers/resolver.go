package application

import "github.com/Senney/todo-app/domain/todo"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoProvider todo.TodoProviderPort
}
