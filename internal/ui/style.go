package ui

import "charm.land/lipgloss/v2"

var (
	Title = lipgloss.NewStyle().
		Background(lipgloss.Color("#8B008B")). // پس‌زمینه بنفش تیره
		Foreground(lipgloss.Color("#FFFFFF")). // متن سفید
		Bold(true)

	Version = lipgloss.NewStyle().
		Background(lipgloss.Color("#3181fa")). // پس‌زمینه صورتی گرم
		Foreground(lipgloss.Color("#1A1A1A")). // متن تیره
		Bold(true)

	Cursor = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F8B55F"))

	Task = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff"))

	Completed = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4DFFBE")).
			Strikethrough(true)

	// header style

	HeaderBg = lipgloss.Color("#3A3A3A")

	StatusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#FF4F7A")).
			Bold(true).
			Padding(0, 1)

	LeftItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#D8D8D8")).
			Background(HeaderBg).
			Padding(0, 1)

	RightItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#A855F7")).
			Padding(0, 1)

	IconStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#7C3AED")).
			Padding(0, 1)
)
