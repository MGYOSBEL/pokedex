package commands

import "github.com/mgyosbel/pokedex/internal/api"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Client                  api.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"map": {
			Name:        "map",
			Description: "Displays the name of 20 location areas in the Pokemon world",
			Callback:    Map,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the name of 20 location areas in the Pokemon world in reverse order",
			Callback:    Mapb,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    Help,
		},
	}
}
