package app

import (
	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	s := tea.NewView("")
	s.AltScreen = true // change cmd view to full view

	if m.width == 0 {
		s.SetContent("Loading...")
		return s
	}

	content := m.header.View(m.width)
	content += "\n\n"
	content += m.todo.View()
	content += "\n\n"
	content += m.help.RenderShort()
	content += "\n\n"
	s.SetContent(content)

	return s
}
