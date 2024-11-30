package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

var (
	pokedex = make(map[string]PokeDetails)
)

type PokeDetails struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
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

func (c *Config) AddPokemonToPokedex(pokemon string, details PokeDetails) error {
	if _, exists := pokedex[pokemon]; exists {
		fmt.Printf("pokemon %s is already stored inside pokedex, hence skipping to add it\n", pokemon)
		return nil
	} else {
		pokedex[pokemon] = details
	}
	return nil
}
