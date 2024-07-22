package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("\nLocation areas:")
	for _, result := range res.Results {
		fmt.Printf(" - %s\n", result.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("\nLocation areas:")
	for _, result := range res.Results {
		fmt.Printf(" - %s\n", result.Name)
	}
	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous
	return nil
}
