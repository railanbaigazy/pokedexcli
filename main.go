package main

import (
	"fmt"
	"log"

	"github.com/railanbaigazy/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()
	res, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	// initCommands()
	// startRepl()
}
