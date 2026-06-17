package todo

import (
	"todo_cli/internal/domain"
	"todo_cli/internal/service"
)

type State struct {
	Cursor  int
	Items   []domain.Todo
	Service *service.TodoService
}

func NewTodoModel(service *service.TodoService) State {
	// TODO: add error and log
	todos, _ := service.GetAll()
	return State{
		Items: todos,
		Service: service,
	}
}

// todo: move to todo service. all of down there 
func (m *State) CursorMoveUp() {
	if m.Cursor > 0 {
		m.Cursor--
	}
}

func (m *State) CursorMoveDown() {
	if m.Cursor < len(m.Items)-1 {
		m.Cursor++
	}
}

func (m *State) DoneToggle() {
	m.Items[m.Cursor].Done = !m.Items[m.Cursor].Done
}
