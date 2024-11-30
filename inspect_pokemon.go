package main

import "fmt"

func (c *Config) inspectPokemon(pokemon string) error {
	if details, exists := pokedex[pokemon]; exists {
		fmt.Printf("Name: %s\n", details.Name)
		fmt.Printf("height: %d\n", details.Height)
		fmt.Printf("Weight: %d\n", details.Weight)
		fmt.Printf("Stats:\n")
		for _, value := range details.Stats {
			fmt.Printf("  -%s: %d\n", value.Stat.Name, value.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, value := range details.Types {
			fmt.Printf("-%s\n", value.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
