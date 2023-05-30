package main

import (
	"log"

	ui "github.com/gizak/termui/v3"

	cmd "github.com/ahmadhabibi14/github-stats-cli/commands"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to terminate termui: %v\n", err)
	}
	defer ui.Close()

	// This is the actual app:
	cmd.Execute()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
