package todo

type Item struct {
	Title string
	Done  bool
}

type Model struct {
	Items  []Item
	Cursor int
}

func NewTodoModel() Model {
	return Model{
		Items: []Item{
			{Title: "Jackal 1"},
			{Title: "Jackal 2"},
			{Title: "Jackal 3"},
			{Title: "Jackal 4"},
		},
	}
}

func (m *Model) CursorMoveUp() {
	if m.Cursor > 0 {
		m.Cursor--
	}
}

func (m *Model) CursorMoveDown() {
	if m.Cursor < len(m.Items)-1 {
		m.Cursor++
	}
}

func (m *Model) DoneToggle() {
	m.Items[m.Cursor].Done = !m.Items[m.Cursor].Done
}
