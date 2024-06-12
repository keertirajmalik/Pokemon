package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/keertirajmalik/pokedexcli/internal/pokeapi"
)

type config struct {
    pokeapiClient pokeapi.Client
    nextLocationsURL *string
    prevLocationURL *string
}

func startRepl(cfg *config) {

    scanner:= bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")

        scanner.Scan()

        words := cleanInput(scanner.Text())
         if len(words) == 0 {
             continue
         }

         commandName := words[0]

        command, exist := getCommands()[commandName]
        if exist {
            err := command.callback(cfg)
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Println("Unknown command")
            continue
        }
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

