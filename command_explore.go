package main

import "fmt"

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no location area provided")
	}
	name := args[0]
	resp, err := cfg.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", resp.Name)
	for _, res := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", res.Pokemon.Name)
	}
	return nil
}
