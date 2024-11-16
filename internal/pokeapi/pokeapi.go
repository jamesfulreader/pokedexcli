package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jamesfulreader/pokedexcli/internal/pokecahe"
)

type Client struct {
	pokeAPIURL string
	httpClient *http.Client
	cache      *pokecahe.Cache
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

func NewClient() *Client {
	return &Client{
		pokeAPIURL: "https://pokeapi.co/api/v2",
		httpClient: &http.Client{},
	}
}

func (c *Client) GetLocations(pageURL *string) (LocationURL, error) {
	url := fmt.Sprintf("%s/location/?offset=0&limit=20", c.pokeAPIURL)
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
		return LocationURL{}, fmt.Errorf("response failed with status code: %d and\nbodye: %s", res.StatusCode, body)
	}

	c.cache.Add(url, body)

	locationURL := LocationURL{}
	err = json.Unmarshal(body, &locationURL)
	if err != nil {
		return LocationURL{}, fmt.Errorf("error unmarshaling json: %s", err)
	}

	return locationURL, nil

}
