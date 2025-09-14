package main

import (
	"time"
	"github.com/MedrekIT/pokedex/internal/pokecache"
	"github.com/MedrekIT/pokedex/internal/pokeapi"
)

func main() {
	conf := &pokeapi.Config{
		Cache: pokecache.NewCache(5 * time.Minute),
	}
	replInit(conf)
}
