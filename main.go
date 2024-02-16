package main

import (
	"fmt"

	"github.com/quinsberry/pokedex/internal/pokeapi"
)

func main() {

	client := pokeapi.NewClient()
	response, err := client.GetLocationAreas()
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	// startRepl()
}
