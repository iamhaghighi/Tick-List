package header

import (
	"todo_cli/internal/styles"

	"charm.land/lipgloss/v2"
)

func (h Model) View(width int) string {
	var leftParts []string
	var rightParts []string

	for _, item := range h.LeftHeader {
		leftParts = append(leftParts, item.Style.Render(item.Text))
	}

	for _, item := range h.RightHeader {
		rightParts = append(rightParts, item.Style.Render(item.Text))
	}

	left := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftParts...,
	)

	right := lipgloss.JoinHorizontal(
		lipgloss.Top,
		rightParts...,
	)

	space := width -
		lipgloss.Width(left) -
		lipgloss.Width(right)

	if space < 0 {
		space = 0
	}

	fill := lipgloss.NewStyle().
		Background(styles.BgHeaderColor).
		Width(space).
		Render("")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		fill,
		right,
	)
}
