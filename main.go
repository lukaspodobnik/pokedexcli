package main

import (
	"time"

	"github.com/lukaspodobnik/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}

	startREPL(cfg)
}
