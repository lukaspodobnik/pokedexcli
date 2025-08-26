package main

import (
	"errors"
	"fmt"
)

func exploreCallback(cfg *config, input ...string) error {
	if len(input) != 1 {
		return errors.New("for explore exactly one argument must be provided: <location-area-name or id>")
	}

	encounters, err := cfg.pokeapiClient.GetPokemonEncounters(input[0])
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Exploring %s...\n", input[0])
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
