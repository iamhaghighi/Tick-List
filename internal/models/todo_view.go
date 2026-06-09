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

	leftItems := []string{
		ui.Title.Render(constants.AppName),
	}
	leftPart := strings.Join(leftItems, "  ") 

	rightItems := []string{
		ui.Version.Render(constants.Version),
		ui.Version.Render(constants.GoVersion),
		ui.Version.Render("Bubbletea v2"),
	}
	rightPart := strings.Join(rightItems, "  ")


	leftWidth := lipgloss.Width(leftPart)
	rightWidth := lipgloss.Width(rightPart)

	padding := m.ts.Width - leftWidth - rightWidth
	if padding < 0 {
		padding = 0
	}

	content := leftPart + strings.Repeat(" ", padding) + rightPart

	return ui.HeaderStyle.Width(m.ts.Width).Render(content)
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
			cursor = ui.Cursor.Render(">")
		}

		checked := "[ ]"
		if m.selected[i] == true {
			checked = ui.Completed.Render("[X]")
			itemStyle = ui.Completed
		}

		mainContent += fmt.Sprintf("%s %s %s\n", cursor, checked, itemStyle.Render(todo))
	}
	s.SetContent(mainContent + "\n" + m.help.RenderShort())
	return s
}
