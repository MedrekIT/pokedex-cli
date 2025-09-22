package main

import (
	"time"
	"github.com/MedrekIT/pokedex-cli/internal/pokecache"
	"github.com/MedrekIT/pokedex-cli/internal/pokeapi"
)

func main() {
	conf := &pokeapi.Config{
		Cache: pokecache.NewCache(5 * time.Minute),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
	replInit(conf)
}
