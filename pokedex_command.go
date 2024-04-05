package main

import (
	"errors"
	"fmt"
)

func commandPokedex(_ *config, args []string, dex *pokedex) error {
	if len(args) > 0 {
		return errors.New("command takes in no arguments")
	}

	if len(dex.pokemens) == 0 {
		fmt.Println("You have no pokemons. You have to catch them all")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range dex.pokemens {
		fmt.Println(" - ", pokemon.Name)
	}
	return nil
}
