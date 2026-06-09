package models

type TodoModel struct {
	todos    []string
	cursor   int
	selected map[int]bool
}

func NewTodoModel() TodoModel {
	return TodoModel{
		todos:    []string{"Season 14", "Notepad", "Think About Future"},
		selected: make(map[int]bool),
	}
}
