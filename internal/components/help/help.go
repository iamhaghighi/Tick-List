package help

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
)

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Save   key.Binding
	A      key.Binding
	E      key.Binding
	Delete key.Binding
	Y      key.Binding
	Esc    key.Binding
	Back   key.Binding
	Quit   key.Binding
}

func NewKeyMap() keyMap {
	return keyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter", "space"),
			key.WithHelp("enter", "select"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc", "b"),
			key.WithHelp("esc/b", "back to menu"),
		),
		A: key.NewBinding(
			key.WithKeys("a", "A"),
			key.WithHelp("a", "add"),
		),
		E: key.NewBinding(
			key.WithKeys("e", "E"),
			key.WithHelp("e", "edit"),
		),
		Delete: key.NewBinding(
			key.WithKeys("delete"),
			key.WithHelp("delete", "delete"),
		),
		Save: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "save"),
		),
		Y: key.NewBinding(
			key.WithKeys("y", "Y"),
			key.WithHelp("y", "confirm"),
		),
		Esc: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("'q'", "quit"),
		),
	}
}

func (k keyMap) MenuKey() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Enter, k.A, k.E, k.Delete, k.Quit}
}

func (k keyMap) EditorKey() []key.Binding {
	return []key.Binding{k.Save, k.Esc}
}

func (k keyMap) DeleteKey() []key.Binding {
	return []key.Binding{k.Y ,k.Esc}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Enter, k.Back},
		{k.Quit},
	}
}

type Model struct {
	Model    help.Model
	ShowHelp bool
	Keys     keyMap
}

func New() Model {
	return Model{
		Model:    help.New(),
		ShowHelp: false,
		Keys:     NewKeyMap(),
	}
}

func (h *Model) SetWidth(width int) {
	h.Model.SetWidth(width)
}

func (h *Model) Toggle() {
	h.ShowHelp = !h.ShowHelp
}

func (h *Model) Hide() {
	h.ShowHelp = false
}

func (h Model) IsVisible() bool {
	return h.ShowHelp
}

func (h Model) MenuRender() string {
	return h.Model.ShortHelpView(h.Keys.MenuKey())
}

func (h Model) EditorKeyRender() string {
	return h.Model.ShortHelpView(h.Keys.EditorKey())
}

func (h Model) DeleteRender() string {
	return h.Model.ShortHelpView(h.Keys.DeleteKey())
}

func (h Model) RenderFull() string {
	return h.Model.FullHelpView(h.Keys.FullHelp())
}
