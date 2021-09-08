package application

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	model "github.com/Senney/todo-app/application/models"
	schema "github.com/Senney/todo-app/application/schema"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	dummyTodo := model.Todo{
		ID: "abc123",
		Content: "This is my todo!",
		Status: model.TodoStatusIncomplete,
	}

	todos = append(todos, &dummyTodo)
	return todos, nil
}

// Query returns schema.QueryResolver implementation.
func (r *Resolver) Query() schema.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
