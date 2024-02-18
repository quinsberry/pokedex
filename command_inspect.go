package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no pokemon name provided")
	}
	name := args[0]
	pokemonMap := cfg.caughtPokemon
	pokemon, ok := pokemonMap[name]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("name: %s\n", pokemon.Name)
	fmt.Printf("height: %d\n", pokemon.Height)
	fmt.Printf("weight: %d\n", pokemon.Weight)
	fmt.Println("stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("types:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	return nil
}
