package main

import (
	"fmt"
	"os"
	"todo_cli/internal/app"
	"todo_cli/internal/constants"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v" || os.Args[1] == "version" || os.Args[1] == "v") {
		fmt.Println("Todo Cli Management " + constants.VERSION)
		return
	}

	p := tea.NewProgram(app.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
