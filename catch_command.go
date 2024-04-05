package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string, dex *pokedex) error {
	if length := len(args); length > 1 || length == 0 {
		return errors.New("command takes in exactly one argument: (name) of the pokemon")
	}
	if _, ok := dex.pokemens[args[0]]; ok {
		fmt.Printf("You already have caught %s\n", args[0])
		return nil
	}
	pokemon, err := cfg.client.CatchPokemon(args[0], cfg.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	if num := rand.Intn(dex.exp); num > pokemon.BaseExperience {
		dex.pokemens[args[0]] = pokemon
		fmt.Printf("%s was caught!\n", args[0])
		dex.exp += pokemon.BaseExperience / 10
	} else {
		fmt.Printf("%s escaped!\n", args[0])
		if dex.exp < pokemon.BaseExperience {
			fmt.Printf("at your current level, you won't be able to catch %s\n", args[0])
		}
	}
	return nil
}
