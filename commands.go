package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(cfg *config, args []string, _ *pokedex) error {
	if len(args) > 0 {
		return errors.New("command currently doesn't take in any arguments")
	}
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		if command.name == "wrong command" {
			continue
		}
		fmt.Println(command.name + ": " + command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, args []string, _ *pokedex) error {
	if len(args) > 0 {
		return errors.New("command currently doesn't take in any arguments")
	}
	os.Exit(0)
	return nil
}

func commandWrong(cfg *config, args []string, _ *pokedex) error {
	fmt.Println("No such command")
	return nil
}
