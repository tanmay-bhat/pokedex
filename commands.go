package main

import (
	"fmt"
	"os"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	nextURL     string
	previousURL string
}

func getCommands() map[string]CliCommand {
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
			callback:    commandMapNext(&config),
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous names of 20 location areas in the Pokemon world",
			callback:    commandMapPrevious(&config),
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
		PokeLocationResponse, err := ListLocations(config.nextURL)
		if err != nil {
			return err
		}
		config.nextURL = PokeLocationResponse.Next
		config.previousURL = PokeLocationResponse.Previous
		for _, location := range PokeLocationResponse.Results {
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
		PokeLocationResponse, err := ListLocations(config.previousURL)
		if err != nil {
			return err
		}
		config.nextURL = PokeLocationResponse.Next
		config.previousURL = PokeLocationResponse.Previous
		for _, location := range PokeLocationResponse.Results {
			fmt.Println(location.Name)
		}
		return nil
	}
}
