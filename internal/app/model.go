package app

import (
	"todo_cli/internal/components/header"
	"todo_cli/internal/components/help"
	"todo_cli/internal/components/todo"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	width  int
	height int

	todo   todo.Model
	help   help.Model
	header header.Model
}

func NewModel() model {
	return model{
		todo:   todo.NewTodoModel(),
		help:   help.NewHelpModel(),
		header: header.NewHeader(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
