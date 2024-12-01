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
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "displays the names of 20 locations",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "displays the previous 20 locations",
			Callback:    CommandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "explore location name given as an arguement",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "catch a pokemon name given as arguemnt",
			Callback:    CommandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "list all of the pokemon in the pokedex",
			Callback:    CommandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "get details on a specific pokemon",
			Callback:    CommandInspect,
		},
	}
}
