package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"time"
	"github.com/Sirkle23/PokedexGo/PokeCache"
)

const (
	BaseURL = "https://pokeapi.co/api/v2/"
)

func cleanInput(text string) []string {
	// Split the input text into words
	words := []string{}
	for _, word := range strings.Split(text, " ") {
		// Convert each word to lowercase and trim spaces
		cleanedWord := strings.Trim(word, ",.!?")
		cleanedWord = strings.ToLower(cleanedWord)
		if cleanedWord != "" {
			words = append(words, cleanedWord)
		}
	}

	// Return the cleaned words
	return words
}

func main() {
	url := BaseURL + "location-area"
	cfg := &config{
		NextURL:     &url,
		PreviousURL: nil,
	}

	cache := PokeCache.NewCache(10 * time.Second)

	// Start REPL (Read-Eval-Print Loop)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		line := scanner.Text()
		
		cleanedWords := cleanInput(line)
		commandFound := false
		for _, command := range getCommands() {
			if cleanedWords[0] == command.name {
				if len(cleanedWords) < 2 {
					cleanedWords = append(cleanedWords, "")
				}

				if err := command.callback(cfg, cache, cleanedWords[1]); err != nil {
					fmt.Printf("Error executing command '%s': %v\n", command.name, err)
				}
				commandFound = true
				break
			}
		}
		if !commandFound && len(cleanedWords) > 0 {
				fmt.Println("Unknown command")
		}
		
		fmt.Print("Pokedex > ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("error while scanning bufio: %v", err)
	}
}