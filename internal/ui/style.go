package ui

import "charm.land/lipgloss/v2"

var (
	//* COLORS
	mainColor   = lipgloss.Color("#BCE29E") // primary brand color
	accentColor = lipgloss.Color("#FF8787") // action / highlight
	MutedColor  = lipgloss.Color("#E5EBB2") // secondary / borders
	BgColor     = lipgloss.Color("#F8C4B4") // background / surfaces

	BlackColor = lipgloss.Color("#222222")
	WhiteColor = lipgloss.Color("#D8D8D8")

	HeaderBgColor = lipgloss.Color("#3A3A3A")

	golangColor = lipgloss.Color("#00ADD8")

	// completed = #4DFFBE
)

var (
	Cursor = lipgloss.NewStyle().
		Foreground(MutedColor)

	Task = lipgloss.NewStyle().
		Foreground(lipgloss.Darken(WhiteColor, 0.3))

	TaskHover = lipgloss.NewStyle().
		Foreground(WhiteColor)

	Completed = lipgloss.NewStyle().
			Foreground(mainColor).
			Strikethrough(true)

	// header style

	StatusStyle = lipgloss.NewStyle().
			Foreground(BlackColor).
			Background(mainColor).
			Bold(true).
			Padding(0, 1)

	LeftItemStyle = lipgloss.NewStyle().
			Foreground(WhiteColor).
			Background(HeaderBgColor).
			Padding(0, 1)

	RightItemStyle = lipgloss.NewStyle().
			Foreground(BlackColor).
			Background(MutedColor).
			Padding(0, 1)

	GolangStyle = lipgloss.NewStyle().
			Foreground(BlackColor).
			Background(golangColor).
			Padding(0, 1)
)
