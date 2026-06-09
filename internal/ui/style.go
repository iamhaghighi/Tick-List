package ui

import "charm.land/lipgloss/v2"

var HeaderStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#f29af5")).
	Height(1)

var Title = lipgloss.NewStyle().
	Background(lipgloss.Color("#8B008B")). // پس‌زمینه بنفش تیره
	Foreground(lipgloss.Color("#FFFFFF")). // متن سفید
	Bold(true).
	Padding(0, 1) // چپ و راست ۱ فاصله

var Version = lipgloss.NewStyle().
	Background(lipgloss.Color("#FF69B4")). // پس‌زمینه صورتی گرم
	Foreground(lipgloss.Color("#1A1A1A")). // متن تیره
	Bold(true).
	Padding(0, 1) // چپ و راست ۱ فاصله

var Cursor = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#F8B55F"))

var Task = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff"))

var Completed = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4DFFBE")).
	Strikethrough(true)
