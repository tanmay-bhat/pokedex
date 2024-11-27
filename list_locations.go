package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

type PokeLocationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func ListLocations(pageURL string) (PokeLocationResponse, error) {
	url := pageURL
	if pageURL == "" {
		url = baseURL + "/location-area"
	}
	resp, err := http.Get(url)
	if err != nil {
		return PokeLocationResponse{}, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeLocationResponse{}, fmt.Errorf("failed to read response body: %v", err)
	}
	if resp.StatusCode > 299 {
		return PokeLocationResponse{}, fmt.Errorf("response failed with status code: %d and body: %s", resp.StatusCode, body)
	}

	pokeLocation := PokeLocationResponse{}
	err = json.Unmarshal(body, &pokeLocation)
	if err != nil {
		return PokeLocationResponse{}, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return pokeLocation, nil
}
