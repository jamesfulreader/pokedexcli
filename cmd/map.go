package cmd

import (
	"fmt"

	"github.com/jamesfulreader/pokedexcli/internal/pokeapi"
)

var (
	nextURL string
	prevURL string
	client  = pokeapi.NewClient()
)

func CommandMap() error {
	fmt.Println()

	var pageURL *string
	if nextURL != "" {
		pageURL = &nextURL
	}

	locationURL, err := client.GetLocations(pageURL)
	if err != nil {
		return err
	}

	if locationURL.Next != nil {
		nextURL = *locationURL.Next
	}
	if locationURL.Previous != nil {
		prevURL = *locationURL.Previous
	}

	for _, locationItem := range locationURL.Results {
		fmt.Println(locationItem.Name)
	}

	fmt.Println()
	return nil
}

func CommandMapB() error {
	if prevURL == "" {
		fmt.Println()
		fmt.Println("cannot go back any further")
		fmt.Println()
		return nil
	}

	var pageURL *string
	if prevURL != "" {
		pageURL = &prevURL
	}

	locationURL, err := client.GetLocations(pageURL)
	if err != nil {
		return err
	}

	if locationURL.Next != nil {
		nextURL = *locationURL.Next
	}
	if locationURL.Previous != nil {
		prevURL = *locationURL.Previous
	}

	for _, locationItem := range locationURL.Results {
		fmt.Println(locationItem.Name)
	}
	fmt.Println()
	return nil
}
