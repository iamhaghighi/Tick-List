package main

import (
	"fmt"
	"os"
	"strings"
	"todo_cli/components"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	todos    []string
	cursor   int
	selected map[int]bool
	help     components.Help
}

func NewModel() model {
	return model{
		todos:    []string{"Season 14", "Notepad", "Think About Future"},
		selected: make(map[int]bool),
		help:     components.NewHelpModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case "enter", "space":
			m.selected[m.cursor] = !m.selected[m.cursor]
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	var titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#49bcff"))

	title := "Todo App"

	width := lipgloss.Width(title) + 4
	if width < 20 {
		width = 20
	}

	top := "╭" + strings.Repeat("-", width-2) + "╮"
	bottom := "╰" + strings.Repeat("─", width-2) + "╯"

	s := top + "\n"
	s += "│" + lipgloss.Place(width-2, 1, lipgloss.Center, lipgloss.Center, titleStyle.Render(title)) + "│\n"
	s += bottom + "\n"

	for i, todo := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if m.selected[i] == true {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, todo)
	}
	return tea.NewView(s + "\n" + m.help.RenderShort() + "\n")
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v" || os.Args[1] == "version") {
		fmt.Println("v1")
		return
	}

	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
