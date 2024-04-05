package main

import (
	"time"

	"github.com/melmustafa/pokedexcli/internal/pokeapi"
	"github.com/melmustafa/pokedexcli/internal/pokecache"
)

func main() {
	config := config{
		client: pokeapi.NewClient(5 * time.Second),
		cache:  pokecache.NewCache(5 * time.Second),
	}
	startRepl(&config)
}
