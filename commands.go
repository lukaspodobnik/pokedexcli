package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, input ...string) error
}

var commands = map[string]cliCommand{}

func init() {
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    exitCallback,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    helpCallback,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next page of location areas.",
		callback:    mapCallback,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous page of location areas.",
		callback:    mapbCallback,
	}

	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Displays all Pokemon located at argument <location-area>",
		callback:    exploreCallback,
	}
}
