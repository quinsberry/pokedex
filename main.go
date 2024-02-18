package main

import (
	"time"

	"github.com/quinsberry/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	exploringLocation   *pokeapi.LocationArea
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
