package main

import (
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var cliCommands map[string]cliCommand

func initCommands() {
	cliCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back",
			callback:    commandMapB,
		},
	}
}

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}
