package main

import (
	"time"

	"github.com/tanmay-bhat/pokedex/internal/cache"
)

const baseURL = "https://pokeapi.co/api/v2"

func main() {
	config := &Config{
		cache: cache.NewCache(5 * time.Minute),
	}
	go config.cache.ReapLoop(5*time.Minute, make(chan bool))
	repl(config)
}
