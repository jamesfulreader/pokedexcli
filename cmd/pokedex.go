package cmd

import "fmt"

func CommandPokedex([]string) error {
	pokemon := config.Pokedex.List()

	if len(pokemon) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}
	fmt.Println("Your caught pokemon")
	for _, entry := range pokemon {
		fmt.Printf("- %s (Base Experience: %d)\n", entry.Name, entry.BaseExperience)
	}
	return nil
}
