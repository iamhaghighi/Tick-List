package models

import (
	"todo_cli/internal/ui/components"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	todo   TodoModel
	help   components.Help
	header components.HeaderModel
	ts     TerminalSizeModel
}

func NewModel() model {
	return model{
		todo:   NewTodoModel(),
		help:   components.NewHelpModel(),
		header: components.NewHeader(),
		ts: NewTerminalSizeModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.ts.Width = msg.Width
		m.ts.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.todo.cursor > 0 {
				m.todo.cursor--
			}
		case "down", "j":
			if m.todo.cursor < len(m.todo.todos)-1 {
				m.todo.cursor++
			}
		case "enter", "space":
			m.todo.selected[m.todo.cursor] = !m.todo.selected[m.todo.cursor]
		}
	}
	return m, nil
}
