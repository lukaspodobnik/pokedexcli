package pokeapi

import "encoding/json"

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
		Base int `json:"base_stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	data, err := c.doCachedGetRequest(url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
