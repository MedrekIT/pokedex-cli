package main

import (
	"errors"
	"fmt"
	"os"
)

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name: "help",
			description: "Displays this help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of the next 20 Pokemon world's locations with each call",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of the previous 20 Pokemon world's locations with each call",
			callback: commandMapb,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandMap(conf *config) error {
	data, err := getLocations(conf.Next)
	if err != nil {
		return err
	}

	conf.Next = data.Next
	conf.Previous = data.Previous

	for _, loc := range data.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}

func commandMapb(conf *config) error {
	if conf.Previous == "" {
		return errors.New("you're on the first page")
	}

	data, err := getLocations(conf.Previous)
	if err != nil {
		return err
	}

	conf.Next = data.Next
	conf.Previous = data.Previous

	for _, loc := range data.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return nil
}

func commandHelp(conf *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	allCommands := getCommands()
	for _, command := range allCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
