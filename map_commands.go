package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args []string, _ *pokedex) error {
	if len(args) > 0 {
		return errors.New("command currently doesn't take in any arguments")
	}
	locations, err := cfg.client.ListLocations(cfg.nextLocationsURL, cfg.cache)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string, _ *pokedex) error {
	if len(args) > 0 {
		return errors.New("command currently doesn't take in any arguments")
	}
	locations, err := cfg.client.ListLocations(cfg.prevLocationsURL, cfg.cache)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
