package main

import (
	"fmt"
	"log"

	"github.com/quinsberry/pokedex/internal/pokeapi"
)

func callbackMapf() error {
	cl := pokeapi.NewClient()
	resp, err := cl.GetLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	for _, res := range resp.Results {
		fmt.Println(res.Name)
	}
	return nil
}

func callbackMapb() error {
	cl := pokeapi.NewClient()
	resp, err := cl.GetLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	for _, res := range resp.Results {
		fmt.Println(res.Name)
	}
	return nil
}
