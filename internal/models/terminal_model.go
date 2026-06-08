package models

type TerminalSizeModel struct {
	Width  int
	Height int
}

func NewTerminalSizeModel() TerminalSizeModel {
	return TerminalSizeModel{
		Width:  0,
		Height: 0,
	}
}
