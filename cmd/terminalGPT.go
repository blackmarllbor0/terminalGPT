package main

import (
	"log"
	"terminalGPT/pkg/ui"
)

func main() {
	userInterface := ui.NewUserInterface()
	if err := userInterface.Start(); err != nil {
		log.Fatal(err)
	}
}
