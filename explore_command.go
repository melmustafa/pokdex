package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string, _ *pokedex) error {
	if length := len(args); length > 1 || length == 0 {
		return errors.New("command takes in exactly one argument: (name) of the area")
	}
	loc, err := cfg.client.ExploreLocation(args[0], cfg.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range loc.Pokemons {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
