package components

import (
	"todo_cli/internal/ui"

	"charm.land/lipgloss/v2"
)

type HeaderItem struct {
	Text  string
	Style lipgloss.Style
}

type HeaderModel struct {
	Left  []HeaderItem
	Right []HeaderItem
}

func NewHeader() HeaderModel {
	return HeaderModel{
		Left: []HeaderItem{
			{
				Text:  "Todo Management",
				Style: ui.StatusStyle,
			},
			{
				Text:  "A Fast, lightweight Cli",
				Style: ui.LeftItemStyle,
			},
		},
		Right: []HeaderItem{
			{
				Text:  "UTF-8",
				Style: ui.RightItemStyle,
			},
			{
				Text:  "⚙",
				Style: ui.IconStyle,
			},
			{
				Text:  "Fish Cake",
				Style: ui.RightItemStyle,
			},
		},
	}
}

func (h HeaderModel) View(width int) string {
	var leftParts []string
	var rightParts []string

	for _, item := range h.Left {
		leftParts = append(leftParts, item.Style.Render(item.Text))
	}

	for _, item := range h.Right {
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
		Background(ui.HeaderBg).
		Width(space).
		Render("")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		fill,
		right,
	)
}
