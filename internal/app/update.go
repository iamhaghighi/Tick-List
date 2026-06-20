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

		if m.activeScreen == TodoScreen {
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit

			case "a", "A":
				m.editor.Input.Placeholder = "write a todo..."
				m.activeScreen = EditorScreen
				m.editor.Mode = editor.CreateMode
				m.editor.Input.SetValue("")
				return m, nil

			case "e", "E":
				if len(m.todoState.Todos) > 0 {
					m.editor.Input.Placeholder = "edit todo..."
					m.activeScreen = EditorScreen
					m.editor.Mode = editor.EditMode
					m.editor.Input.SetValue(m.todoState.Todos[m.todoState.Cursor].Title)
					return m, nil
				}

			case "delete":
				if len(m.todoState.Todos) > 0 {
					m.editor.Input.Placeholder = "  y - confirm to delete"
					m.activeScreen = EditorScreen
					m.editor.Mode = editor.DeleteMode
					m.editor.Input.SetValue("")
					return m, nil
				}

			case "enter":
				if len(m.todoState.Todos) > 0 {
					id := m.todoState.Todos[m.todoState.Cursor].ID
					done := m.todoState.Todos[m.todoState.Cursor].Done
					_ = m.todoState.Service.Toggle(Ctx, id, !done)
					m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
				}
				return m, nil

			default:
				m.todoState.Update(msg)
				return m, nil
			}
		}

		if m.activeScreen == EditorScreen {
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit

			case "esc":
				m.activeScreen = TodoScreen
				return m, nil

			case "enter":
				value := m.editor.Input.Value()

				if m.editor.Mode == editor.CreateMode {
					if value != "" {
						_, _ = m.todoState.Service.Create(Ctx, value)
						m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
						m.activeScreen = TodoScreen
						return m, nil
					}
				} else if m.editor.Mode == editor.EditMode {
					if value != "" && len(m.todoState.Todos) > 0 {
						_, _ = m.todoState.Service.Update(Ctx, domain.Todo{
							ID:    m.todoState.Todos[m.todoState.Cursor].ID,
							Title: value,
						})
						m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
						m.activeScreen = TodoScreen
						return m, nil
					}
				} else if m.editor.Mode == editor.DeleteMode {
					if value == "y" && len(m.todoState.Todos) > 0 {
						id := m.todoState.Todos[m.todoState.Cursor].ID
						_ = m.todoState.Service.Delete(Ctx, id)
						m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
						m.activeScreen = TodoScreen
						return m, nil
					} else if value == "n" || value == "N" {
						m.activeScreen = TodoScreen
						return m, nil
					}
				}
				return m, nil

			default:
				m.editor.Input, cmd = m.editor.Input.Update(msg)
				return m, cmd
			}
		}
	}

	return m, cmd
}
