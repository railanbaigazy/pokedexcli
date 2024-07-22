package main

import (
	"fmt"
	"log"

	"github.com/railanbaigazy/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	pokeapiClient := pokeapi.NewClient()
	res, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nLocation areas:")
	for _, result := range res.Results {
		fmt.Printf(" - %s\n", result.Name)
	}
	return nil
}
