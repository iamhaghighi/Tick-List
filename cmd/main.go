package main

import (
	"fmt"
	"os"
	"todo_cli/internal/models"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v" || os.Args[1] == "version") {
		fmt.Println("v1")
		return
	}

	p := tea.NewProgram(models.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
