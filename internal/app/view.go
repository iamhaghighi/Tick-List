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

	mainContent := m.header.View(m.width)
	mainContent += "\n\n"
	mainContent += m.todo.View()

	
	s.SetContent(mainContent + "\n" + m.help.RenderShort())
	return s
}
