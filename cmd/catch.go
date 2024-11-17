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
	data, err := config.Client.GetPokemon(&name)
	if err != nil {
		return err
	}

	pokemonBaseXP := data.BaseExperience
	catchRate := 50.0 / float64(pokemonBaseXP)
	randomNumber := rand.Float64()

	if randomNumber <= catchRate {
		fmt.Println("You caught", data.Name)
	} else {
		fmt.Println("Failed to catch", data.Name)
	}
	return nil
}
