package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mgyosbel/pokedex/internal/commands"
)

const prompt = "Pokedex > "

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := commands.NewConfig()

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
			continue
		}
		arguments := text[1:]
		cfg.CommandArguments = arguments
		cmd.Callback(cfg)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
