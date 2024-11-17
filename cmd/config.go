package cmd

import "github.com/jamesfulreader/pokedexcli/internal/pokeapi"

type Config struct {
	Client  *pokeapi.Client
	NextURL string
	PrevURL string
}

var config = &Config{
	Client: pokeapi.NewClient(),
}
