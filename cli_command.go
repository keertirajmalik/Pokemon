package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get Next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get Previous  page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commndExplore,
		},
		"catch": {
			name:        "catch <Pokemon_name>",
			description: "Catch the pokemon found in area",
			callback:    commandCatch,
		},
        "inspect": {
            name: "inspect <Pokemon_name>",
            description: "View details about a caught pokemon",
            callback: commandInspect,
        },
        "pokedex": {
            name: "pokedex",
            description: "See all the pokemon you have caught",
            callback: commandPokedex,
        },
	}
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commndExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")

	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chance := rand.Intn(pokemon.BaseExperience)

	if chance > 40 {
		fmt.Printf("%s escaped\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s is caught\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

func commandInspect(cfg *config, args ...string) error {
    if len(args)   != 1 {
        return errors.New("You must provide pokemon name")
    }

    name := args[0]
    pokemon, ok := cfg.caughtPokemon[name]

    if !ok {
        return errors.New("You have not caught that pokemon")
    }

    fmt.Println("Name:", pokemon.Name)
    fmt.Println("Height:", pokemon.Height)
    fmt.Println("Weight:", pokemon.Weight)
    fmt.Println("Stats:")

    for _, stat := range pokemon.Stats {
        fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Println("Types:")
    for _, typeInfo := range pokemon.Types {
        fmt.Println(" -",typeInfo.Type.Name)
    }
    return nil
}

func commandPokedex(cfg *config, args ...string) error {

    if len(cfg.caughtPokemon) == 0 {
        fmt.Println("No pokemon caught")
    }

    fmt.Println("Your Pokedex")
    for _, pokemon := range cfg.caughtPokemon {
        fmt.Println(" - ", pokemon.Name)
    }
    return nil
}
