package main

import (
	"log"
	"fmt"
	"os"
	"encoding/json"
	"math/rand"
	"github.com/Sirkle23/PokedexGo/PokedexAPI"
	"github.com/Sirkle23/PokedexGo/PokeCache"
)

type config struct {
	NextURL 		*string
	PreviousURL 	*string
}

func CommandExit(cfg *config, cache *PokeCache.Cache, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *config, cache *PokeCache.Cache, parameter string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func CommandMapf(cfg *config, cache *PokeCache.Cache, parameter string) error {
	body, exists := cache.Get(*cfg.NextURL)
	if !exists {
		body = PokedexAPI.GetPokedexBytes(*cfg.NextURL)
		cache.Add(*cfg.NextURL, body)
	}

	var data PokedexAPI.Map_strct
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	// Update command configuration
	cfg.NextURL = data.Next
	cfg.PreviousURL = data.Previous

	return nil
}

func CommandMapb(cfg *config, cache *PokeCache.Cache, parameter string) error {
	if cfg.PreviousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	body, exists := cache.Get(*cfg.PreviousURL)
	if !exists {
		body = PokedexAPI.GetPokedexBytes(*cfg.PreviousURL)
		cache.Add(*cfg.PreviousURL, body)
	}

	var data PokedexAPI.Map_strct
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	// Update command configuration
	cfg.NextURL = data.Next
	cfg.PreviousURL = data.Previous

	return nil
}

func CommandExplore(cfg *config, cache *PokeCache.Cache, parameter string) error {
	if parameter == "" {
		fmt.Println("Please provide a location to explore.")
		return nil
	}

	exploreURL := fmt.Sprintf("%slocation-area/%s", BaseURL, parameter)

	body, exists := cache.Get(exploreURL)
	if !exists {
		body = PokedexAPI.GetPokedexBytes(exploreURL)
		cache.Add(exploreURL, body)
	}

	var data PokedexAPI.Expl_strct
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, pokemonEncounter := range data.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func CommandCatch(cfg *config, cache *PokeCache.Cache, parameter string) error {
	if parameter == "" {
		fmt.Println("Please provide a pokemon to catch.")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", parameter)

	catchURL := fmt.Sprintf("%spokemon/%s", BaseURL, parameter)

	body, exists := cache.Get(catchURL)
	if !exists {
		body = PokedexAPI.GetPokedexBytes(catchURL)
		cache.Add(catchURL, body)
	}

	var data PokedexAPI.Pokemon_strct
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	caught := false
	random_number := rand.Intn(100)
	fmt.Println(random_number, "is the random number generated.")
	fmt.Println("Base Experience of", parameter, "is", data.BaseExperience)
	if data.BaseExperience/2 > 75 {
		if random_number > 75 {
			caught = true
		}
	} else {
		if random_number > data.BaseExperience/2 {
			caught = true
		}
	}
	
	if caught {
		cache.AddPokemon(parameter, data)
		fmt.Println(parameter, "was caught!")
	} else {
		fmt.Println(parameter, " escaped!")
	}

	return nil
}

func CommandInspect(cfg *config, cache *PokeCache.Cache, parameter string) error {
	if parameter == "" {
		fmt.Println("Please provide a pokemon to inspect.")
		return nil
	}

	pokemon, exists := cache.GetPokemon(parameter)
	if !exists {
		fmt.Printf("You have not caught %s.\n", parameter)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, pokeStat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", pokeStat.Stat.Name, pokeStat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("\t- %s\n", pokeType.Type.Name)
	}

	return nil
}

func CommandPokedex(cfg *config, cache *PokeCache.Cache, parameter string) error {
	if len(cache.CaughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemon yet.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cache.CaughtPokemon {
		fmt.Printf("\t- %s\n", pokemon.Name)
	}

	return nil
}
