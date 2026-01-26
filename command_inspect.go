package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Printf("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, exists := cfg.caughtPokemon[name]
	if exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats: \n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t-%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types: \n")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("\t - %s \n", pokeType.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon yet")
	}
	return nil
}
