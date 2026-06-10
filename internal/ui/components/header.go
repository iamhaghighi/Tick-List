package components

import (
	"todo_cli/internal/constants"
	"todo_cli/internal/ui"

	"charm.land/lipgloss/v2"
)

type HeaderItem struct {
	Text  string
	Style lipgloss.Style
}

type HeaderModel struct {
	LeftHeader  []HeaderItem
	RightHeader []HeaderItem
}

func NewHeader() HeaderModel {
	return HeaderModel{
		LeftHeader: []HeaderItem{
			{
				Text:  constants.APP_NAME,
				Style: ui.TitleStyle,
			},
			{
				Text:  constants.DESCRIPTION,
				Style: ui.DescriptionStyle,
			},
		},
		RightHeader: []HeaderItem{
			{
				Text:  constants.GO_VERSION,
				Style: ui.GolangStyle,
			},
			{
				Text:  constants.VERSION,
				Style: ui.VersionStyle,
			},
		},
	}
}

func (h HeaderModel) View(width int) string {
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
		Background(ui.BgHeaderColor).
		Width(space).
		Render("")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		fill,
		right,
	)
}
