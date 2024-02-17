package main

import (
	"time"

	"github.com/quinsberry/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute),
	}
	startRepl(&cfg)
}
