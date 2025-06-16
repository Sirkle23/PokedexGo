package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		line := scanner.Text()
		
		cleanedWords := cleanInput(line)
		fmt.Printf("Your command was: %s\n", cleanedWords[0])
		
		fmt.Print("Pokedex > ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("error while scanning bufio: %v", err)
	}
}

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