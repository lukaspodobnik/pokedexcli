package main

import (
	"errors"
	"fmt"
)

func pokedexCallback(cfg *config, input ...string) error {
	if len(input) != 0 {
		return errors.New("this command does not accept any arguments")
	}

	fmt.Println()
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	fmt.Println()

	return nil
}
