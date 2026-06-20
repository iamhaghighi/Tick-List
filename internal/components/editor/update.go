package editor

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) tea.Cmd{
	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)
	return cmd
}
