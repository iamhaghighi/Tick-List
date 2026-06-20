package app

import (
	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	s := tea.NewView("")
	s.AltScreen = true

	if m.width == 0 {
		s.SetContent("Loading...")
		return s
	}

	con := m.header.View(m.width)
	con += "\n\n"

	switch m.activeScreen {

	case TodoScreen:
		con += m.todoState.View()

	case EditorScreen:
		todoTitle := m.todoState.Todos[m.todoState.Cursor].Title
		con += m.editor.View(todoTitle)

	}

	s.SetContent(con)

	return s
}
