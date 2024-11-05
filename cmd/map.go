package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationURL struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var nextURL string
var prevURL string
var count int = 0
var pokeapi string = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"

func CommandMap() error {
	fmt.Println()
	if count >= 0 {
		if nextURL != "" { // Check if nextURL is not nil
			pokeapi = nextURL
		}
	}

	res, err := http.Get(pokeapi)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	locationURL := LocationURL{}
	err = json.Unmarshal(body, &locationURL)
	if err != nil {
		log.Fatalf("error with unmarshaling JSON: %s", err)
	}

	if locationURL.Next != nil {
		nextURL = *locationURL.Next
	}

	if locationURL.Previous != nil {
		prevURL = *locationURL.Previous
	}

	locationsArray := locationURL.Results
	for _, locationItem := range locationsArray {
		fmt.Println(locationItem.Name)
	}
	fmt.Println()
	count++
	return nil
}

func CommandMapB() error {
	fmt.Println()
	if count <= 1 {
		fmt.Println()
		fmt.Println("cannot go back any further")
		fmt.Println()
		count = 0
	}
	if prevURL != "" { // Check if prevURL is not nil
		pokeapi = prevURL
	}

	res, err := http.Get(pokeapi)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	locationURL := LocationURL{}
	err = json.Unmarshal(body, &locationURL)
	if err != nil {
		log.Fatalf("error with unmarshaling JSON: %s", err)
	}

	if locationURL.Next != nil {
		nextURL = *locationURL.Next
	}

	if locationURL.Previous != nil {
		prevURL = *locationURL.Previous
	}

	locationsArray := locationURL.Results
	for _, locationItem := range locationsArray {
		fmt.Println(locationItem.Name)
	}
	fmt.Println()
	count--
	return nil
}
