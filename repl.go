package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		var text string
		if scanner.Scan() {
			text = scanner.Text()
		} else {
			if scanner.Err() != nil {
				fmt.Printf("Could not capture input. Got error: %v\n", scanner.Err())
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

		err := command.callback()
		if err != nil {
			fmt.Printf("Callback of command did not succeed: %v\n", err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
