package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/melmustafa/pokedexcli/internal/pokecache"
)

func (c Client) ExploreLocation(areaName string, cache pokecache.Cache) (RespLocation, error) {
	url := baseURL + "/location-area/" + areaName

	if data, ok := cache.Get(url); ok {
		return handleRawArea(data)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return RespLocation{}, err
	}

	if resp.StatusCode > 299 {
		return RespLocation{}, errors.New("location not found")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	go cache.Add(url, data)
	if err != nil {
		return RespLocation{}, err
	}

	return handleRawArea(data)
}

func handleRawArea(rawData []byte) (RespLocation, error) {
	result := RespLocation{}
	if err := json.Unmarshal(rawData, &result); err != nil {
		return RespLocation{}, err
	}
	return result, nil
}
