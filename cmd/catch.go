package cmd

import (
	"errors"
	"fmt"
	"math/rand"
)

func CommandCatch(pokemonName []string) error {
	if len(pokemonName) <= 1 {
		return errors.New("please provide a pokemon name")
	}

	name := pokemonName[1]
	pokemon, err := config.Client.GetPokemon(&name)
	if err != nil {
		return err
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokemon.Name)

	pokemonBaseXP := pokemon.BaseExperience
	catchRate := 50.0 / float64(pokemonBaseXP)
	randomNumber := rand.Float64()

	if randomNumber <= catchRate {
		config.Pokedex.Add(pokemon.Name, pokemon)
		fmt.Printf("You caught %s!\n", pokemon.Name)
	} else {
		fmt.Printf("%s got away!\n", pokemon.Name)
	}
	return nil
}
