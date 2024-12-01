package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jamesfulreader/pokedexcli/internal/models"
	"github.com/jamesfulreader/pokedexcli/internal/pokecache"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	cache      *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		baseURL:    "https://pokeapi.co/api/v2",
		httpClient: &http.Client{},
		cache:      pokecache.NewCache(5 * time.Minute),
	}
}

func (c *Client) makeRequest(url string, target interface{}) error {
	// Check cache first
	if data, found := c.cache.Get(url); found {
		return json.Unmarshal(data, target)
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("API request failed with status %d: %s", res.StatusCode, body)
	}

	c.cache.Add(url, body)

	return json.Unmarshal(body, target)
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

func (c *Client) GetLocations(pageURL *string) (LocationURL, error) {
	url := fmt.Sprintf("%s/location-area/?offset=0&limit=20", c.baseURL)
	if pageURL != nil {
		url = *pageURL
	}

	var locations LocationURL
	err := c.makeRequest(url, &locations)
	return locations, err
}

type LocationArea struct {
	Name              string                    `json:"name"`
	PokemonEncounters []models.PokemonEncounter `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(locationName *string) (LocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%s", c.baseURL, *locationName)

	var location LocationArea
	err := c.makeRequest(url, &location)
	return location, err
}

func (c *Client) GetPokemon(pokemonName *string) (models.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", c.baseURL, *pokemonName)

	var pokemon models.Pokemon
	err := c.makeRequest(url, &pokemon)
	return pokemon, err
}
