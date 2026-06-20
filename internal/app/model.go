package app

import (
	"context"
	"todo_cli/internal/components/editor"
	"todo_cli/internal/components/header"
	"todo_cli/internal/components/help"
	"todo_cli/internal/components/todo"
	"todo_cli/internal/service"

	tea "charm.land/bubbletea/v2"
)

var Ctx = context.Background()

type model struct {
	width        int
	height       int
	activeScreen Screen

	todoState todo.State
	help      help.Model
	header    header.Model
	editor    editor.Model
}

func NewModel(service *service.TodoService) model {
	return model{
		activeScreen: TodoScreen,
		todoState:    todo.New(service),
		help:         help.New(),
		header:       header.NewHeader(),
		editor:       editor.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
