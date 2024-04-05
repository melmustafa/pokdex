package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/melmustafa/pokedexcli/internal/pokecache"
)

func (c Client) CatchPokemon(pokemonName string, cache pokecache.Cache) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if data, ok := cache.Get(url); ok {
		return handleRawPokemon(data)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return RespPokemon{}, err
	}

	if resp.StatusCode > 299 {
		return RespPokemon{}, errors.New("no such pokemon")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	go cache.Add(url, data)
	if err != nil {
		return RespPokemon{}, err
	}

	return handleRawPokemon(data)
}

func handleRawPokemon(rawData []byte) (RespPokemon, error) {
	result := RespPokemon{}
	if err := json.Unmarshal(rawData, &result); err != nil {
		return RespPokemon{}, err
	}
	return result, nil
}
