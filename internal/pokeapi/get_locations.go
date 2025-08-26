package pokeapi

import (
	"encoding/json"
)

type RespLocationArea struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

func (c *Client) GetLocations(nextURL *string) (RespLocationArea, error) {
	url := baseURL + "/location-area"
	if nextURL != nil {
		url = *nextURL
	}

	data, err := c.doCachedGetRequest(url)
	if err != nil {
		return RespLocationArea{}, err
	}

	locationArea := RespLocationArea{}
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return RespLocationArea{}, err
	}

	return locationArea, nil
}
