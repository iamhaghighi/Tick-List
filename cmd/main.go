package main

import (
	"fmt"
	"log"
	"os"
	"todo_cli/internal/app"
	"todo_cli/internal/constants"
	"todo_cli/internal/repository/sqlite"
	"todo_cli/internal/service"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v" || os.Args[1] == "version" || os.Args[1] == "v") {
		fmt.Println("Todo Cli Management " + constants.VERSION)
		return
	}

	repo, err := sqlite.New(`..\internal\storage\notes.db`)
	if err != nil {
		log.Fatal(err)
	}

	service := service.NewTodoService(repo)

	p := tea.NewProgram(app.NewModel(service))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
