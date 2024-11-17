package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jamesfulreader/pokedexcli/internal/pokecache"
)

type Client struct {
	pokeAPIURL string
	httpClient *http.Client
	cache      *pokecache.Cache
}

type LocationURL struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Will need something like:
type LocationArea struct {
	Name              string             // "pastoria-city-area"
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

// Each Pokemon in the area
type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"` // nesting the Pokemon info
}

// The actual Pokemon details
type Pokemon struct {
	Name string `json:"name"` // "tentacool"
	URL  string `json:"url"`  // not needed but usually good to include for completeness
}

func NewClient() *Client {
	return &Client{
		pokeAPIURL: "https://pokeapi.co/api/v2",
		httpClient: &http.Client{},
		cache:      pokecache.NewCache(5 * time.Minute),
	}
}

func (c *Client) GetLocations(pageURL *string) (LocationURL, error) {
	url := fmt.Sprintf("%s/location-area/?offset=0&limit=20", c.pokeAPIURL)
	if pageURL != nil {
		url = *pageURL
	}

	if data, found := c.cache.Get(url); found {
		locationURL := LocationURL{}
		err := json.Unmarshal(data, &locationURL)
		if err != nil {
			return LocationURL{}, fmt.Errorf("error unmarshaling cached json: %s", err)
		}
		return locationURL, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationURL{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationURL{}, err
	}

	if res.StatusCode > 299 {
		return LocationURL{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	c.cache.Add(url, body)

	locationURL := LocationURL{}
	err = json.Unmarshal(body, &locationURL)
	if err != nil {
		return LocationURL{}, fmt.Errorf("error unmarshaling json: %s", err)
	}

	return locationURL, nil
}

func (c *Client) GetLocationArea(locationName *string) (LocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%s", c.pokeAPIURL, *locationName)

	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error with get %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationDetails := LocationArea{}
	err = json.Unmarshal(body, &locationDetails)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error unmarshaling loc area %s", err)
	}
	return locationDetails, nil
}
