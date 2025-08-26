package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catchCallback(cfg *config, input ...string) error {
	if len(input) != 1 {
		return errors.New("for explore exactly one argument must be provided: <location-area-name or id>")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(input[0])
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if rand.Intn(pokemon.BaseExperience) > 40 {
		fmt.Printf("%s was caught!", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!", pokemon.Name)
	}
	fmt.Println()

	return nil
}
