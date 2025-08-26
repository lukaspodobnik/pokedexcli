package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/lukaspodobnik/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) doCachedGetRequest(url string) ([]byte, error) {
	data, hit := c.cache.Get(url)
	if !hit {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []byte{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return []byte{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}

		c.cache.Add(url, data)
	}

	return data, nil
}
