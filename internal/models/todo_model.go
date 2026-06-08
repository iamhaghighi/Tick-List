package models

import (
	"todo_cli/internal/ui/components"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	todos    []string
	cursor   int
	selected map[int]bool
	help     components.Help
	ts       TerminalSizeModel
}

func NewModel() model {
	return model{
		todos:    []string{"Season 14", "Notepad", "Think About Future"},
		selected: make(map[int]bool),
		help:     components.NewHelpModel(),
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
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case "enter", "space":
			m.selected[m.cursor] = !m.selected[m.cursor]
		}
	}
	return m, nil
}
