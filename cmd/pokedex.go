package cmd

import (
	"errors"
	"fmt"
)

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

func CommandInspect(pokemonName []string) error {
	if len(pokemonName) <= 1 {
		return errors.New("please enter a pokemon name")
	}

	pokemon, exists := config.Pokedex.Inspect(pokemonName[1])
	if !exists {
		fmt.Println("You haven't caught that pokemon")
		return nil
	}
	fmt.Printf("\nName: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf(" -%s\n", typeInfo.Type.Name)
	}
	fmt.Println()
	return nil
}
