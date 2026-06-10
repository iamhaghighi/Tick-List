package header

import (
	"todo_cli/internal/constants"
	"todo_cli/internal/styles"

	"charm.land/lipgloss/v2"
)

type HeaderItem struct {
	Text  string
	Style lipgloss.Style
}

type Model struct {
	LeftHeader  []HeaderItem
	RightHeader []HeaderItem
}

func NewHeader() Model {
	return Model{
		LeftHeader: []HeaderItem{
			{
				Text:  constants.APP_NAME,
				Style: styles.TitleStyle,
			},
			{
				Text:  constants.DESCRIPTION,
				Style: styles.DescriptionStyle,
			},
		},
		RightHeader: []HeaderItem{
			{
				Text:  constants.GO_VERSION,
				Style: styles.GolangStyle,
			},
			{
				Text:  constants.VERSION,
				Style: styles.VersionStyle,
			},
		},
	}
}
