package editor

import "fmt"

func (m Model) View(todoTitle string) string {
	title := ""

	if m.Mode == CreateMode {
		title = "Create Todo"
	}

	if m.Mode == EditMode {
		title = "Edit: " + todoTitle
	}

	if m.Mode == DeleteMode {
		title = fmt.Sprintf("Delete this todo? (%s) y/N", todoTitle)
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n", title, m.Input.View(),
	)
}
