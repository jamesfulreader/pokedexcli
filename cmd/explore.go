package cmd

import (
	"fmt"
)

func CommandExplore(location []string) error {
	fmt.Println("Exploring ", location[1])
	url := location[1]
	data, err := config.Client.GetLocationArea(&url)
	if err != nil {
		return err
	}
	fmt.Printf("You find %d pokemon\n", len(data.PokemonEncounters))
	for _, pokemon := range data.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
