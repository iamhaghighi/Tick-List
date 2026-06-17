package todo

import tea "charm.land/bubbletea/v2"

func (m *State) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			m.CursorMoveUp()

		case "down", "j":
			m.CursorMoveDown()

		case "enter":
			m.DoneToggle()
		}
	}
}
