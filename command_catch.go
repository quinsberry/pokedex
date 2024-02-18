package main

import (
	"fmt"
	"math/rand"

	"github.com/quinsberry/pokedex/internal/pokeapi"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no pokemon name provided")
	}
	name := args[0]

	if cfg.exploringLocation == nil || !isPokemonInLocation(name, cfg.exploringLocation) {
		return fmt.Errorf("%s is not in the current location", name)
	}

	if _, ok := cfg.caughtPokemon[name]; ok {
		return fmt.Errorf("%s is already in your pokedex", name)
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", name)
	}

	cfg.caughtPokemon[name] = pokemon
	fmt.Printf("%s was caught!\n", name)
	return nil
}

func isPokemonInLocation(pokemonName string, location *pokeapi.LocationArea) bool {
	for _, res := range location.PokemonEncounters {
		if res.Pokemon.Name == pokemonName {
			return true
		}
	}
	return false
}
