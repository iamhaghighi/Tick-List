package app

import (
	"context"
	"log"
	"todo_cli/internal/components/editor"
	"todo_cli/internal/domain"

	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	m.editor.Input, cmd = m.editor.Input.Update(msg)
	cmds = append(cmds, cmd)

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
			if m.activeScreen == TodoScreen {
				m.editor.Input.Placeholder = "edit todo..."

				m.activeScreen = EditorScreen
				m.editor.Mode = editor.EditMode

				m.editor.Input.SetValue("")
			}
		case "delete":
			if m.activeScreen == TodoScreen {
				m.editor.Input.Placeholder = "y/N"

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
				id := m.todoState.Todos[m.todoState.Cursor].ID
				done := m.todoState.Todos[m.todoState.Cursor].Done
				//todo: error handling
				m.todoState.Service.Toggle(Ctx, id, !done)
				m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)
			}

			if m.activeScreen == EditorScreen {

				if m.editor.Mode == editor.CreateMode {
					value := m.editor.Input.Value()

					if value != "" {
						m.todoState.Service.Create(Ctx, value)
						todos, _ := m.todoState.Service.GetAll(context.Background())
						m.todoState.Todos = todos

						m.activeScreen = TodoScreen
					}
				}

				if m.editor.Mode == editor.EditMode {
					value := m.editor.Input.Value()

					if value != "" {

						m.todoState.Service.Update(Ctx, domain.Todo{
							ID:    m.todoState.Todos[m.todoState.Cursor].ID,
							Title: value,
						})

						m.todoState.Todos, _ = m.todoState.Service.GetAll(Ctx)

						m.activeScreen = TodoScreen
					}
				}

				if m.editor.Mode == editor.DeleteMode {
					value := m.editor.Input.Value()
					switch value {
					case "y":
						id := m.todoState.Todos[m.todoState.Cursor].ID
						if err := m.todoState.Service.Delete(Ctx, id); err != nil {
							log.Fatalf("unable to delete todo: %v", err)
						}
						var err error
						m.todoState.Todos, err = m.todoState.Service.GetAll(Ctx)
						if err != nil {
							log.Fatalf("unable to get todos: %v", err)
						}
						m.activeScreen = TodoScreen
					case "n", "N":
						m.activeScreen = TodoScreen
					}
				}
			}

		}
	}

	switch m.activeScreen {
	case TodoScreen:
		m.todoState.Update(msg)
	case EditorScreen:
		cmd = m.editor.Update(msg)
	}

	return m, cmd
}
