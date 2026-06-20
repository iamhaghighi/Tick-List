package editor

import "fmt"

func (m Model) View() string {
	title := ""
	help := ""

	if m.Mode == CreateMode {
		title = "Create Todo:"
		help = m.help.EditorKeyRender()
	}

	if m.Mode == EditMode {
		title = "Edit Todo: "
		help = m.help.EditorKeyRender()
	}

	if m.Mode == DeleteMode {
		title = "Delete Todo?"
		help = m.help.DeleteRender()

	}

	return fmt.Sprintf(
		"%s\n%s\n\n%s", title, m.Input.View(), help,
	)
}
