package cmd

type CliCommand struct {
	Name        string
	Description string
	Callback    func([]string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a list of all available commands",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Shows the next 20 locations in the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous 20 locations",
			Callback:    CommandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "Lists all Pokemon that can be found in a specific location",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempts to catch a Pokemon",
			Callback:    CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Lists all Pokemon you have caught",
			Callback:    CommandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Shows detailed information about a caught Pokemon",
			Callback:    CommandInspect,
		},
	}
}
