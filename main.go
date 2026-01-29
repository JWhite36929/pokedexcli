package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jwhite36929/pokedexcli/internal/pokeapi"
)

func main() {
	savePath := "pokedexSave.json"
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	fmt.Println("Attmepting to read save data.")
	if _, err := os.Stat(savePath); err == nil {
		content, err := os.ReadFile(savePath)
		if err != nil {
			log.Fatal(err)
		}

		var saveData configSave
		err = json.Unmarshal(content, &saveData)
		if err != nil {
			log.Fatal(err)
		}
		cfg := &config{
			caughtPokemon: saveData.CaughtPokemon,
			pokeapiClient: pokeClient,
		}
		startRepl(cfg)

	} else {
		fmt.Println("No save file found. Creating a new save...")
		cfg := &config{
			caughtPokemon: map[string]pokeapi.Pokemon{},
			pokeapiClient: pokeClient,
		}
		startRepl(cfg)
	}

}
