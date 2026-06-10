package styles

import "charm.land/lipgloss/v2"

var (
	TextPrimaryColor   = lipgloss.Color(Light)
	TextSecondaryColor = lipgloss.Color(Dark)

	BgTitleColor  = lipgloss.Color(primary)
	BgHeaderColor = lipgloss.Color(Medium)

	StatusCompleted = lipgloss.Color(primary)
	CursorColor     = lipgloss.Color(secondary)

	GolangColor = lipgloss.Color(Golang)
)

var (
	// view style
	Cursor = lipgloss.NewStyle().
		Foreground(CursorColor)

	Task = lipgloss.NewStyle().
		Foreground(lipgloss.Darken(TextPrimaryColor, 0.3))

	TaskHover = lipgloss.NewStyle().
			Foreground(TextPrimaryColor)

	Completed = lipgloss.NewStyle().
			Foreground(BgTitleColor).
			Strikethrough(true)

	// header style
	TitleStyle = lipgloss.NewStyle().
			Foreground(TextSecondaryColor).
			Background(BgTitleColor).
			Bold(true).
			Padding(0, 1)

	DescriptionStyle = lipgloss.NewStyle().
				Foreground(TextPrimaryColor).
				Background(BgHeaderColor).
				Padding(0, 1)

	VersionStyle = lipgloss.NewStyle().
			Foreground(TextSecondaryColor).
			Background(CursorColor).
			Padding(0, 1)

	GolangStyle = lipgloss.NewStyle().
			Foreground(TextSecondaryColor).
			Background(GolangColor).
			Padding(0, 1)
)
