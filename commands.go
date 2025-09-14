package main

import (
	"errors"
	"fmt"
	"os"
	"github.com/MedrekIT/pokedex/internal/pokeapi"
)

type CliCommand struct {
	name string
	description string
	callback func(*pokeapi.Config, []string) error
}

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
		"explore": {
			name: "explore <area_name>",
			description: "Displays all Pokemons located in provided area",
			callback: commandExplore,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}

func commandExit(conf *pokeapi.Config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandMap(conf *pokeapi.Config, params []string) error {
	data, err := pokeapi.GetLocations(conf.Next, conf)
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

func commandMapb(conf *pokeapi.Config, params []string) error {
	if conf.Previous == "" {
		return errors.New("you're on the first page")
	}

	data, err := pokeapi.GetLocations(conf.Previous, conf)
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

func commandExplore(conf *pokeapi.Config, params []string) error {
	if len(params) != 1 {
		return errors.New("incorrect usage!\nTry 'explore <area-name>'")
	}

	data, err := pokeapi.GetPokemons(params[0], conf)
	if err != nil {
		return err
	}
	
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", params[0])
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandHelp(conf *pokeapi.Config, params []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	allCommands := getCommands()
	for _, command := range allCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
