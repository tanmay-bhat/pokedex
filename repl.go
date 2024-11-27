package main

import (
	"fmt"
	"maps"
	"slices"
)

var config Config

func repl() {
	var input string
	// infinite for loop to keep the terminal running
	for {
		fmt.Print("Pokedex > ")
		fmt.Scanln(&input)

		command, exists := getCommands()[input]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			commands := slices.Collect(maps.Keys(getCommands()))
			fmt.Printf("Unknown command specified, available commands are: %v\n", commands)
		}
	}
}
