package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{}

func initCommands() {
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    exit,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    help,
	}
}
