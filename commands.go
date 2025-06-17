package main

import (
	"github.com/Sirkle23/PokedexGo/PokeCache"
)

type cliCommand struct {
	name				string
	description			string
	callback			func(*config, *PokeCache.Cache, string) error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Display names of next 20 locations",
			callback:    CommandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display names of previous 20 locations",
			callback:    CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all pokemon at location",
			callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display information about a caught pokemon",
			callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all caught pokemon",
			callback:    CommandPokedex,
		},
	}
}