package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
    name string
    description string
    callback func(*config) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Display a help message",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "map": {
            name: "map",
            description: "Get Next page of locations",
            callback: commandMapf,
        },
        "mapb": {
            name: "mapb",
            description: "Get Previous  page of locations",
            callback: commandMapb,
        },
    }
}

func commandHelp(cfg *config) error{
    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, cmd := range getCommands() {
        fmt.Printf("%s: %s\n", cmd.name,cmd.description)
    }
    fmt.Println()
    return nil
}

func commandExit(cfg *config) error {
    os.Exit(0)
    return nil
}

func commandMapf(cfg *config) error {
    locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL  = locationResp.Next
    cfg.prevLocationURL = locationResp.Previous

    for _, loc := range locationResp.Results {
        fmt.Println(loc.Name)
    }
    return nil
}

func commandMapb(cfg *config) error {
    if cfg.prevLocationURL == nil {
        return errors.New("you're on the first page")
    }
    locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL  = locationResp.Next
    cfg.prevLocationURL = locationResp.Previous

    for _, loc := range locationResp.Results {
        fmt.Println(loc.Name)
    }
    return nil
}
