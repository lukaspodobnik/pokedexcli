package main

import (
	"errors"
	"fmt"
)

func exploreCallback(cfg *config, input ...string) error {
	if len(input) != 1 {
		return errors.New("for explore exactly one argument must be provided: <location-area>")
	}

	encounters, err := cfg.pokeapiClient.GetPokemonEncounters(input[0])
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Exploring %s...\n", input[0])
	for _, encounter := range encounters.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
