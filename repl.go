package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

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
            err := command.callback()
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

func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Display a help message",
            callback: commandHelp,
        },
        "exit": {
            name: "help",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
    }
}

