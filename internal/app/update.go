package app

import (
	"todo_cli/internal/components/editor"
	"todo_cli/internal/domain"

	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		key := msg.String()

		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "a", "A":
			if m.activeScreen == TodoScreen {
				m.editor.Input.Placeholder = "write a todo..."
				m.activeScreen = EditorScreen
				m.editor.Mode = editor.CreateMode
				m.editor.Input.SetValue("")
			}

		case "e", "E":
			if m.activeScreen == TodoScreen && len(m.todoState.Todos) > 0 {
				m.editor.Input.Placeholder = "edit todo..."
				m.activeScreen = EditorScreen
				m.editor.Mode = editor.EditMode
				m.editor.Input.SetValue(m.todoState.Todos[m.todoState.Cursor].Title)
			}

		case "delete":
			if m.activeScreen == TodoScreen && len(m.todoState.Todos) > 0 {
				m.editor.Input.Placeholder = "  y - confirm to delete"
				m.activeScreen = EditorScreen
				m.editor.Mode = editor.DeleteMode
				m.editor.Input.SetValue("")
			}

		case "esc":
			if m.activeScreen == EditorScreen {
				m.activeScreen = TodoScreen
			}

		case "enter":
			if m.activeScreen == TodoScreen {
				if len(m.todoState.Todos) > 0 {
					id := m.todoState.Todos[m.todoState.Cursor].ID
					done := m.todoState.Todos[m.todoState.Cursor].Done
					_ = m.todoState.Service.Toggle(Ctx, id, !done)
					m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
				}
			}

			if m.activeScreen == EditorScreen {
				value := m.editor.Input.Value()

				if m.editor.Mode == editor.CreateMode {
					if value != "" {
						_, _ = m.todoState.Service.Create(Ctx, value)
						m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
						m.activeScreen = TodoScreen
					}
				} else if m.editor.Mode == editor.EditMode {
					if value != "" && len(m.todoState.Todos) > 0 {
						_, _ = m.todoState.Service.Update(Ctx, domain.Todo{
							ID:    m.todoState.Todos[m.todoState.Cursor].ID,
							Title: value,
						})
						m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
						m.activeScreen = TodoScreen
					}
				} else if m.editor.Mode == editor.DeleteMode {
					if value == "y" && len(m.todoState.Todos) > 0 {
						id := m.todoState.Todos[m.todoState.Cursor].ID
						_ = m.todoState.Service.Delete(Ctx, id)
						m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
						m.activeScreen = TodoScreen
					} else if value == "n" || value == "N" {
						m.activeScreen = TodoScreen
					}
				}
			}
		}

		if m.activeScreen == EditorScreen {
			m.editor.Input, cmd = m.editor.Input.Update(msg)
		}
	}

	if m.activeScreen == TodoScreen {
		m.todoState.Update(msg)
	}

	return m, cmd
}
