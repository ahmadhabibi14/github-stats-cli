package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	cmd "github.com/ahmadhabibi14/github-stats-cli/commands"
)

func main() {
	// This is the actual app, yeahhh:
	cmd.Execute() // <<<<< main app okay??
	// until this

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM)
	<-terminate
	fmt.Println("Exiting")
}
