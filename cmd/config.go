package cmd

import (
	"github.com/jamesfulreader/pokedexcli/internal/pokeapi"
	"github.com/jamesfulreader/pokedexcli/internal/pokedex"
)

type Config struct {
	Client  *pokeapi.Client
	NextURL string
	PrevURL string
	Pokedex *pokedex.Pokedex
}

var config = &Config{
	Client:  pokeapi.NewClient(),
	Pokedex: pokedex.New(),
}
