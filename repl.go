package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/MedrekIT/pokedex/internal/pokeapi"
)


func cleanInput(test string) []string {
	slicedText := strings.Fields(strings.ToLower(test))
	return slicedText
}

func replInit(conf *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			allCommands := getCommands()
			if command, ok := allCommands[input[0]]; ok {
				err := command.callback(conf, input[1:])
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
