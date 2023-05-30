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

	// This is the actual app, yeahhh:
	cmd.Execute() // <<<<< main app okay??
	// until this

	// terminate := make(chan os.Signal, 1)
	// signal.Notify(terminate, os.Interrupt, syscall.SIGTERM)
	// <-terminate
	// fmt.Println("Exiting")

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
