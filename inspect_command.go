package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string, dex *pokedex) error {
	if length := len(args); length > 1 || length == 0 {
		return errors.New("command takes in exactly one argument: (name) of the pokemon")
	}
	if len(dex.pokemens) == 0 {
		fmt.Println("You haven't caught any pokemons. Catch some first. Catch them all")
		return nil
	}

	pokemon, ok := dex.pokemens[args[0]]
	if !ok {
		fmt.Println("You have caught this pokemon it. Try to catch it first.")
		return nil
	}
	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, Type := range pokemon.Types {
		fmt.Println(" - ", Type.Type.Name)
	}
	return nil
}
