package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("your pokedex is empty")
	}
	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
