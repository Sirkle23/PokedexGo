package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

var cliCommands map[string]cliCommand

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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func main() {
	// Initialize commands map here
    cliCommands = make(map[string]cliCommand)

    // Add commands to the map
    cliCommands["exit"] = cliCommand{ 
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
    cliCommands["help"] = cliCommand{ 
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	// Start REPL (Read-Eval-Print Loop)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		line := scanner.Text()
		
		cleanedWords := cleanInput(line)
		commandFound := false
		for _, command := range cliCommands {
			if cleanedWords[0] == command.name {
				if err := command.callback(); err != nil {
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