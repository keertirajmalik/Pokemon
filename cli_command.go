package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
    name string
    description string
    callback func() error
}

func commandHelp() error{
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

func commandExit() error {
    os.Exit(0)
    return nil
}
