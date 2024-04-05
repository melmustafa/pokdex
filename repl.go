package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/melmustafa/pokedexcli/internal/pokeapi"
	"github.com/melmustafa/pokedexcli/internal/pokecache"
)

type config struct {
	cache            pokecache.Cache
	client           pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type pokedex struct {
	pokemens map[string]pokeapi.RespPokemon
	exp      int
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	dex := pokedex{
		pokemens: make(map[string]pokeapi.RespPokemon),
		exp:      100,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		args := handleInput(scanner.Text())
		cmd := args[0]
		handleRepl(cmd, cfg, args[1:], dex)
	}

}

func handleRepl(cmd string, cfg *config, args []string, dex pokedex) {
	if len(cmd) == 0 {
		return
	}
	command, ok := getCommands()[cmd]
	if !ok {
		command = getCommands()["wrong command"]
	}
	err := command.callback(cfg, args, &dex)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		fmt.Println()
	}
}

func handleInput(line string) []string {
	output := strings.ToLower(line)
	lineWords := strings.Fields(output)
	return lineWords
}
