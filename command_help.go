package main

import "fmt"

func commandHelp(cfg *config) error {
    fmt.Println()
    fmt.Println("Welcome to Pokedex")
    fmt.Println("Usage:")
    fmt.Println("Type 'exit' to quit")
    fmt.Println()
    for _, cmd := range getCommands() {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    fmt.Println()
    return nil
}
