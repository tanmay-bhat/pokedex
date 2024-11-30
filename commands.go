package main

import (
	"fmt"
	"os"

	"github.com/tanmay-bhat/pokedex/internal/cache"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	nextURL     string
	previousURL string
	cache       *cache.Cache
}

func getCommands(config *Config) map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next names of 20 location areas in the Pokemon world",
			callback:    commandMapNext(config),
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous names of 20 location areas in the Pokemon world",
			callback:    commandMapPrevious(config),
		},
		"explore": {
			name:        "explore",
			description: "Displays the list of all the Pokémon in a given area",
			callback:    nil,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the specified Pokemon",
			callback:    nil,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemom for its abilities",
			callback:    nil,
		},
	}
}

func commandHelp() error {
	helpText := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
    `
	fmt.Println(helpText)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMapNext(config *Config) func() error {
	return func() error {
		resp, err := config.ListLocations(config.nextURL)
		if err != nil {
			return err
		}
		config.nextURL = resp.Next
		config.previousURL = resp.Previous
		for _, location := range resp.Results {
			fmt.Println(location.Name)
		}

		return nil
	}
}

func commandMapPrevious(config *Config) func() error {
	return func() error {
		if config.previousURL == "" {
			fmt.Println("Cannot go to previous page, you are at the first page")
			return nil
		}
		resp, err := config.ListLocations(config.previousURL)
		if err != nil {
			return err
		}
		config.nextURL = resp.Next
		config.previousURL = resp.Previous
		for _, location := range resp.Results {
			fmt.Println(location.Name)
		}
		return nil
	}
}

func commandMapExplore(config *Config, location string) func() error {
	return func() error {
		fmt.Printf("Exploring %s...\n", location)
		resp, err := config.ExploreLocation(location)
		if err != nil {
			return err
		}
		fmt.Printf("Found Pokemon:\n")
		for _, pokemon := range resp.PokemonEncounters {
			fmt.Println(pokemon.Pokemon.Name)
		}
		return nil
	}
}

func commandCatch(config *Config, pokemon string) func() error {
	return func() error {
		fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
		pokeDetails, err := config.GetPokemonDetails(pokemon)
		if err != nil {
			return err
		}
		caught, err := config.CatchPokemon(pokeDetails)
		if err != nil {
			return err
		}
		if !caught {
			fmt.Println("pikachu escaped!")
		} else {
			fmt.Println("pikachu was caught!")
			config.AddPokemonToPokedex(pokemon, pokeDetails)
		}
		return nil
	}
}

func commandInspect(config *Config, pokemon string) func() error {
	return func() error {
		config.inspectPokemon(pokemon)
		return nil
	}
}
