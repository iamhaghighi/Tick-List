package todo

import (
	"context"
	"todo_cli/internal/components/help"
	"todo_cli/internal/domain"
	"todo_cli/internal/service"
)

type State struct {
	Cursor  int
	Todos   []domain.Todo
	Service *service.TodoService

	help help.Model
}

func New(service *service.TodoService) State {
	// TODO: add error and log
	ctx := context.Background()
	todos, _ := service.GetAll(ctx)
	return State{
		Todos:   todos,
		Service: service,
		help:    help.New(),
	}
}

// todo: move to todo service
func (m *State) CursorMoveUp() {
	if m.Cursor > 0 {
		m.Cursor--
	}
}

func (m *State) CursorMoveDown() {
	if m.Cursor < len(m.Todos)-1 {
		m.Cursor++
	}
}

// func (m *State) DoneToggle() {
// 	m.Todos[m.Cursor].Done = !m.Todos[m.Cursor].Done
// }
