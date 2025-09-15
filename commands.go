package main

import (
	"errors"
	"fmt"
	"os"
	"github.com/MedrekIT/pokedex/internal/pokeapi"
	"math/rand"
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
		"catch": {
			name: "catch <pokemon_name>",
			description: "Allows trying to catch a Pokemon! If successful, adds it to your Pokedex",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon_name>",
			description: "Displays caught Pokemon's statistics",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays all caught Pokemon",
			callback: commandPokedex,
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
		return errors.New("incorrect usage!\nTry 'explore <area_name>'")
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

func commandCatch(conf *pokeapi.Config, params []string) error {
	if len(params) != 1 {
		return errors.New("incorrect usage!\nTry 'catch <pokemon_name>'")
	}

	data, err := pokeapi.GetPokemon(params[0], conf)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])
	chance := rand.Intn(data.BaseExperience)
	if chance < 40 {
		conf.Pokedex[params[0]] = data
		fmt.Printf("%s was caught!\n", params[0])
		fmt.Printf("It's added to your Pokedex and you may now inspect it with the 'inspect' command.\n")
	} else {
		fmt.Printf("%s escaped!\n", params[0])
	}

	return nil
}

func commandInspect(conf *pokeapi.Config, params[]string) error {
	if len(params) != 1 {
		return errors.New("incorrect usage!\nTry 'inspect <pokemon_name>'")
	}

	pokemon, ok := conf.Pokedex[params[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Heigh: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokeType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokeType.Type.Name)
	}

	return nil
}

func commandPokedex(conf *pokeapi.Config, params []string) error {
	if len(conf.Pokedex) == 0 {
		fmt.Printf("You haven't caught any Pokemon yet!\n")
		return nil
	}

	fmt.Printf("Your Pokedex:\n")
	for pokemon, _ := range conf.Pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}

func commandHelp(conf *pokeapi.Config, params []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("'%s': %s\n", command.name, command.description)
	}

	return nil
}
