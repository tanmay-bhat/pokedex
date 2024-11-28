package main

import (
	"time"

	"github.com/tanmay-bhat/pokedex/internal/cache"
)

func main() {
	config := &Config{
		cache: cache.NewCache(5 * time.Minute),
	}
	go config.cache.ReapLoop(5*time.Minute, make(chan bool))
	repl(config)
}
