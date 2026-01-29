package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jwhite36929/pokedexcli/internal/pokeapi"
)

type configSave struct {
	CaughtPokemon map[string]pokeapi.Pokemon `json:"caughtPokemon"`
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Saving data...")

	saveData := configSave{
		CaughtPokemon: cfg.caughtPokemon,
	}

	jsonData, err := json.MarshalIndent(saveData, "", " ")
	if err != nil {
		log.Fatal("Error marshalling to JSON:", err)
	}

	saveFile := "pokedexSave.json"
	err = os.WriteFile(saveFile, jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing to file: ", err)
	}

	fmt.Println("Data saved successfully, Goodbye!")
	os.Exit(0)
	return nil
}
