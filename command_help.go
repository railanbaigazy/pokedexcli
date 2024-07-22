package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range cliCommands {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Printf("\n")
	return nil
}
