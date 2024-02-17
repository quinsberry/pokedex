package main

import "github.com/quinsberry/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
