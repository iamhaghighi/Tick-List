package editor

import (
	"todo_cli/internal/components/help"

	"charm.land/bubbles/v2/textinput"
)

type Mode int

const (
	CreateMode Mode = iota
	EditMode
	DeleteMode
)

type Model struct {
	Input  textinput.Model
	Mode   Mode
	TodoID string

	help help.Model
}

func New() Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 120
	ti.SetWidth(40)

	return Model{
		Input: ti,
		help:  help.New(),
	}
}
