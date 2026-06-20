package editor

import "fmt"

func (m Model) View(todoTitle string) string {
	title := ""
	help := ""

	if m.Mode == CreateMode {
		title = "Create Todo"
		help = m.help.EditorKeyRender()
	}

	if m.Mode == EditMode {
		title = "Edit: " + todoTitle
		help = m.help.EditorKeyRender()
	}

	if m.Mode == DeleteMode {
		title = fmt.Sprintf("Delete this todo? (%s)", todoTitle)
		help = m.help.DeleteRender()

	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s", title, m.Input.View(), help,
	)
}
