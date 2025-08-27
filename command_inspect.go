package main

import (
	"errors"
	"fmt"
)

func inspectCallback(cfg *config, input ...string) error {
	if len(input) != 1 {
		return errors.New("for inspect exactly one argument must be provided: <pokemon>")
	}

	pokemon, exits := cfg.caughtPokemon[input[0]]
	if !exits {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println()
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.Base)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	fmt.Println()

	return nil
}
