package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	application "github.com/Senney/todo-app/application/resolvers"
	"github.com/Senney/todo-app/application/schema"
	"github.com/Senney/todo-app/domain/entities"
	"github.com/Senney/todo-app/domain/todo"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	todos := [...]entities.Todo{
		{ID: "abc", Content: "Todo #1!", Status: entities.TodoStatusCompleted},
		{ID: "def", Content: "Todo #2!", Status: entities.TodoStatusIncomplete},
	}

	todoProvider := todo.TodoProviderAdapter{
		Todos: todos[:],
	}

	resolver := application.Resolver{
		TodoProvider: &todoProvider,
	}

	srv := handler.NewDefaultServer(schema.NewExecutableSchema(schema.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
