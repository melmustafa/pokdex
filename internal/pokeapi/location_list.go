package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/melmustafa/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, cache pokecache.Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := cache.Get(url); ok {
		return handleRawLocationsData(data)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	go cache.Add(url, data)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return handleRawLocationsData(data)
}

func handleRawLocationsData(rawData []byte) (RespShallowLocations, error) {
	locationsResp := RespShallowLocations{}
	if err := json.Unmarshal(rawData, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}
	return locationsResp, nil
}
