// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package application

import (
	"fmt"
	"io"
	"strconv"
)

type Todo struct {
	ID      string     `json:"id"`
	Content string     `json:"content"`
	Status  TodoStatus `json:"status"`
}

type TodoStatus string

const (
	TodoStatusComplete   TodoStatus = "COMPLETE"
	TodoStatusIncomplete TodoStatus = "INCOMPLETE"
)

var AllTodoStatus = []TodoStatus{
	TodoStatusComplete,
	TodoStatusIncomplete,
}

func (e TodoStatus) IsValid() bool {
	switch e {
	case TodoStatusComplete, TodoStatusIncomplete:
		return true
	}
	return false
}

func (e TodoStatus) String() string {
	return string(e)
}

func (e *TodoStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoStatus", str)
	}
	return nil
}

func (e TodoStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
