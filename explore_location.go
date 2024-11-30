package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Config) ExploreLocation(location string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area/" + location

	// Check the cache for the current nextURL
	if cachedLocation, found, err := c.cache.Get(url); found {
		if err != nil {
			return LocationAreaResponse{}, err
		}
		resp := LocationAreaResponse{}
		if err := json.Unmarshal(cachedLocation, &resp); err != nil {
			return LocationAreaResponse{}, err
		}
		return resp, nil
	}
	resp, err := http.Get(url)

	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to fetch Pokemons in specified Area: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to read response body: %v", err)
	}
	if resp.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("response failed with status code: %d and body: %s", resp.StatusCode, body)
	}

	pokeArea := LocationAreaResponse{}
	err = json.Unmarshal(body, &pokeArea)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to parse JSON: %v", err)
	}

	c.cache.Add(url, body)
	return pokeArea, nil
}
