package main

import "fmt"

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no location area provided")
	}
	name := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationArea(name)
	if err != nil {
		return err
	}

	cfg.exploringLocation = &locationArea
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, res := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", res.Pokemon.Name)
	}
	return nil
}
