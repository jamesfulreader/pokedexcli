package cmd

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/jamesfulreader/pokedexcli/internal/pokedex"
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
		fmt.Printf("You caught %s!\n", data.Name)
		entry := pokedex.PokedexEntry{
			Name:           data.Name,
			URL:            data.URL,
			BaseExperience: data.BaseExperience,
		}
		config.Pokedex.Add(data.Name, entry)
	} else {
		fmt.Printf("%s got away!\n", data.Name)
	}
	return nil
}
