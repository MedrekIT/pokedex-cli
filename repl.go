package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)


func cleanInput(test string) []string {
	slicedText := strings.Fields(strings.ToLower(test))
	return slicedText
}

func replInit() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			allCommands := getCommands()
			if command, ok := allCommands[input[0]]; ok {
				err := command.callback()
				if err != nil {
					fmt.Errorf("%w", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
