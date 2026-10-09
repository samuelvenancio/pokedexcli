package main

import (
	"time"

	"github.com/samuelvenancio/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
