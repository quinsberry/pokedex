package main

import (
	"errors"
	"fmt"
)

func callbackMapf(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}

	for _, res := range resp.Results {
		fmt.Println(res.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaUrl == nil {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}

	for _, res := range resp.Results {
		fmt.Println(res.Name)
	}
	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
