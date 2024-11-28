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
		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
