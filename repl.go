package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lukaspodobnik/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		var text string
		if scanner.Scan() {
			text = scanner.Text()
		} else {
			if scanner.Err() != nil {
				fmt.Println(scanner.Err())
			} else {
				fmt.Println("Close the Pokedex due to reading EOF")
			}
			break
		}

		input := cleanInput(text)
		if len(input) == 0 {
			continue
		}

		command, ok := commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
