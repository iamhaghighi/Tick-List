package app

import (
	"todo_cli/internal/components/header"
	"todo_cli/internal/components/help"
	"todo_cli/internal/components/todo"
	"todo_cli/internal/service"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	width  int
	height int

	todo   todo.State
	help   help.Model
	header header.Model
}

func NewModel(service *service.TodoService) model {
	return model{
		todo:   todo.NewTodoModel(service),
		help:   help.NewHelpModel(),
		header: header.NewHeader(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
