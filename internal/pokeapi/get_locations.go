package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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

	data, hit := c.cache.Get(url)
	if !hit {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocationArea{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocationArea{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespLocationArea{}, err
		}

		c.cache.Add(url, data)
	}

	locationArea := RespLocationArea{}
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return RespLocationArea{}, err
	}

	return locationArea, nil
}
