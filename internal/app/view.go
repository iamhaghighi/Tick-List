package app

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	s := tea.NewView("")
	s.AltScreen = true // change cmd view to full view

	if m.width == 0 {
		s.SetContent("Loading...")
		return s
	}

	// content := m.header.View(m.width)
	// content += "\n\n"
	// content += m.todo.View()
	// content += "\n\n"
	// content += m.help.RenderShort()
	// content += "\n\n"

	con := lipgloss.JoinVertical(
		lipgloss.Left,
		m.header.View(m.width),
		"",
		m.todo.View(),
		"",
		m.help.RenderShort(),
	)

	s.SetContent(con)

	return s
}
