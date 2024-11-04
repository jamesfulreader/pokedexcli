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
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// var locationMeta []LocationURL
// var locationMetadata = make([]any, 0)

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

	// locationMeta = append(locationMeta, locationURL)
	// locationMetadata = append(locationMetadata, locationURL.Count)
	// locationMetadata = append(locationMetadata, locationURL.Next)
	// locationMetadata = append(locationMetadata, locationURL.Previous)

	locationsArray := locationURL.Results
	for _, locationItem := range locationsArray {
		fmt.Println(locationItem.Name)
	}

	return nil
}

func CommandMapB() error {
	// prev := LocationURL{}
	// err = json.Unmarshal(locationMetadata[0], &prev)
	// if err != nil {
	// 	log.Fatalf("error with unmarshaling JSON: %s", err)
	// }
	// fmt.Println(locationMetadata[0])
	// fmt.Println(locationMetadata[1])
	// fmt.Println(locationMetadata[2])
	// if locationMetadata == string {
	// 	fmt.Println("cannot go back")
	// }

	// fmt.Println(reflect.TypeOf(locationMetadata[0]))
	return nil
}
