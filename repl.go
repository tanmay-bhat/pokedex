package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(config *Config) {
	commands := getCommands(config)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		command := parts[0]
		args := parts[1:]

		if cmd, found := commands[command]; !found {
			fmt.Println("Unknown command. Type 'help' for a list of commands")
		} else if command == "explore" {
			if len(args) < 1 {
				fmt.Println("Error: location argument is required")
			}
			location := string(args[0])

			exploreCallback := commandMapExplore(config, location)
			if err := exploreCallback(); err != nil {
				fmt.Printf("Error running commandMapExplore: %v\n", err)
			}
		} else if err := cmd.callback(); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
