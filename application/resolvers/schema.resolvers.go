package application

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	application "github.com/Senney/todo-app/application/models"
	application1 "github.com/Senney/todo-app/application/schema"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*application.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns application1.QueryResolver implementation.
func (r *Resolver) Query() application1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
