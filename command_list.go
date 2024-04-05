package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string, *pokedex) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"explore": {
			name:        "explore",
			description: "Displays a help message",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"wrong command": {
			name:        "wrong command",
			description: "No such command",
			callback:    commandWrong,
		},
		"map": {
			name:        "map",
			description: "Display 20 locations in the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations in the pokemon world",
			callback:    commandMapb,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokeman",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display the status of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all the caught pokemons",
			callback:    commandPokedex,
		},
	}
}
