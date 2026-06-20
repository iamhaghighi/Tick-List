package messages

import "todo_cli/internal/domain"

type TodoCreatedMsg struct {
	Todo domain.Todo
}

type TodoDeletedMsg struct {
	ID string
}

type TodoUpdatedMsg struct {
	Todo domain.Todo
}