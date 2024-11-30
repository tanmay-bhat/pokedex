package main

import "fmt"

func (c *Config) Pokedex() {
	if len(pokedex) == 0 {
		fmt.Println("You have not caught any pokemon yet.")
	} else {
		fmt.Println("Your Pokedex:")
		for pokemon := range pokedex {
			fmt.Printf(" - %s\n", pokemon)
		}
	}
}
