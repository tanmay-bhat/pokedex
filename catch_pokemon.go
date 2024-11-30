package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type PokeDetails struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func (c *Config) GetPokemonDetails(pokemon string) (PokeDetails, error) {
	//make the API call and get info about the pokemon
	url := baseURL + "/pokemon/" + pokemon

	if cachedPokemon, found, err := c.cache.Get(url); found {
		if err != nil {
			return PokeDetails{}, err
		}
		resp := PokeDetails{}
		if err := json.Unmarshal(cachedPokemon, &resp); err != nil {
			return PokeDetails{}, err
		}
		return resp, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return PokeDetails{}, fmt.Errorf("failed to fetch Pokemons in specified Area: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeDetails{}, fmt.Errorf("failed to read response body: %v", err)
	}
	if resp.StatusCode > 299 {
		return PokeDetails{}, fmt.Errorf("response failed with status code: %d and body: %s", resp.StatusCode, body)
	}

	pokedetails := PokeDetails{}
	err = json.Unmarshal(body, &pokedetails)
	if err != nil {
		return PokeDetails{}, fmt.Errorf("failed to parse JSON: %v", err)
	}
	c.cache.Add(url, body)
	return pokedetails, nil
}

func (c *Config) CatchPokemon(p PokeDetails) (caught bool, err error) {
	baseExperience := p.BaseExperience
	//The higher the base experience, the harder it should be to catch.
	// legendary pokemons with exp > 200, if chance is baseExperience > 100, we can say we caught it
	// rare pokemons with exp > 100 but < 200, if baseExperience > 60, we can say we caught it.

	chance := rand.Intn(baseExperience)

	switch {
	case baseExperience > 200:
		caught = chance > baseExperience-100
	case baseExperience > 100:
		caught = chance > baseExperience-40
	default:
		caught = chance > baseExperience-10
	}
	return caught, nil
}

func (c *Config) AddPokemonToPokedex(p Pokemon)
