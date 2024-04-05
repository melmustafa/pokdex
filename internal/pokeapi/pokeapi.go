package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type RespLocation struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type RespPokemon struct {
	Id             int    `json:"id"`
	BaseExperience int    `json:"base_experience"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
