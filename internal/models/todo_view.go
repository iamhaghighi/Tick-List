package models

import (
	"fmt"
	"strings"
	"todo_cli/internal/constants"
	"todo_cli/internal/ui"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) headerView() string {
	if m.ts.Width == 0 {
		return ""
	}

	padding := m.ts.Width - lipgloss.Width(constants.AppName) - lipgloss.Width(constants.Version)
	if padding < 0 {
		padding = 0
	}
	
	content := constants.AppName + strings.Repeat(" ", padding) + constants.Version
	return ui.HeaderStyle.Render(content)
}

func (m model) View() tea.View {
	s := tea.NewView("")
	s.AltScreen = true

	if m.ts.Width == 0 {
		s.SetContent("Loading...")
		return s
	}

	mainContent := m.headerView() + "\n\n"

	for i, todo := range m.todos {

		itemStyle := ui.Task
		cursor := " "

		if m.cursor == i {
			cursor = ui.Cursor.Render("▶")
		}

		checked := "[ ]"
		if m.selected[i] == true {
			checked = ui.Completed.Render("[✓]")
			itemStyle = ui.Completed
		}

		mainContent += fmt.Sprintf("%s %s %s\n", cursor, checked, itemStyle.Render(todo))
	}
	s.SetContent(mainContent + "\n" + m.help.RenderShort())
	return s
}
