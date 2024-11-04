package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationURL struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap() error {

	res, err := http.Get("https://pokeapi.co/api/v2/location/?offset=0&limit=20")
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

	locationsArray := locationURL.Results
	for _, locationItem := range locationsArray {
		fmt.Println(locationItem.Name)
	}

	return nil
}
