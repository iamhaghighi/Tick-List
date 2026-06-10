package app

import tea "charm.land/bubbletea/v2"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m.todo.CursorMoveUp()
		case "down", "j":
			m.todo.CursorMoveDown()
		case "enter", "space":
			m.todo.DoneToggle()
		}
	}
	return m, nil
}
