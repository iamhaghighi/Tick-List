package todo

import (
	"fmt"
	"todo_cli/internal/styles"
)

func (m Model) View() string {
	mainContent := ""
	for i, todo := range m.Items {

		itemStyle := styles.Task
		cursor := " "

		if m.Cursor == i {
			itemStyle = styles.TaskHover
			cursor = styles.Cursor.Render("▶")
		}

		checked := "[ ]"
		if m.Items[i].Done == true {
			checked = styles.Completed.UnsetStrikethrough().Render("[✓]")
			itemStyle = styles.Completed
		}

		mainContent += fmt.Sprintf("%s %s %s\n", cursor, checked, itemStyle.Render(todo.Title))
	}
	return mainContent
}
