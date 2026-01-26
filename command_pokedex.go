package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	fmt.Printf("Your Pokedex currently has: \n")

	for k := range cfg.caughtPokemon {
		fmt.Printf("\t- %s\n", k)
	}

	return nil
}
