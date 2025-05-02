package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mgyosbel/pokedex/commands"
)

const prompt = "Pokedex > "

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var cfg commands.Config

	for {
		fmt.Print(prompt)

		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		availableCommands := commands.GetCommands()

		cmd, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
		}

		cmd.Callback(&cfg)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
