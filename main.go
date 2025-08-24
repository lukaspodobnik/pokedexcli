package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
				fmt.Printf("Close the Pokedex due to reading EOF.\n")
			}
			break
		}

		input := cleanInput(text)
		fmt.Printf("Your command was: %s\n", input[0])
	}
}
