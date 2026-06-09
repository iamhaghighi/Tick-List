package models

import (
	"fmt"
	"todo_cli/internal/ui"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	s := tea.NewView("")
	s.AltScreen = true

	if m.ts.Width == 0 {
		s.SetContent("Loading...")
		return s
	}

	mainContent := m.header.View(m.ts.Width)
	mainContent += "\n\n"

	for i, todo := range m.todo.todos {

		itemStyle := ui.Task
		cursor := " "

		if m.todo.cursor == i {
			cursor = ui.Cursor.Render(">")
		}

		checked := "[ ]"
		if m.todo.selected[i] == true {
			checked = ui.Completed.Render("[X]")
			itemStyle = ui.Completed
		}

		mainContent += fmt.Sprintf("%s %s %s\n", cursor, checked, itemStyle.Render(todo))
	}
	s.SetContent(mainContent + "\n" + m.help.RenderShort())
	return s
}
