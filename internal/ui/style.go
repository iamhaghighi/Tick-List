package ui

import "charm.land/lipgloss/v2"

var HeaderStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#9af5d3")). // رنگ آبی ویندوز
	Foreground(lipgloss.Color("#000")).
	Bold(true).
	Height(1)

var Title = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#000")).
	Background(lipgloss.Color("#4DFFBE")).
	Bold(true)

var Cursor = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#F8B55F"))

var Task = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff"))

var Completed = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4DFFBE"))
