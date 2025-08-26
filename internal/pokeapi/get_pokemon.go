package pokeapi

import (
	"encoding/json"
)

type RespPokemonEncounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetPokemonEncounters(locationArea string) (RespPokemonEncounters, error) {
	url := baseURL + "/location-area/" + locationArea

	data, err := c.doCachedGetRequest(url)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	encounters := RespPokemonEncounters{}
	if err := json.Unmarshal(data, &encounters); err != nil {
		return RespPokemonEncounters{}, err
	}

	return encounters, nil
}
