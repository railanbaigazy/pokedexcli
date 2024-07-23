package main

import (
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var cliCommands map[string]cliCommand

func initCommands() {
	cliCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
		"explore": {
			name:        "explore {location_area}",
			description: "List the pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Get information about your pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See your pokemon in the pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
