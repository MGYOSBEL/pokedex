package commands

import "github.com/mgyosbel/pokedex/internal/api"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Client                  *api.Client
	NextLocationAreaUrl     *string
	PreviousLocationAreaUrl *string
	CommandArguments        []string
	UserPokedex             map[string]api.Pokemon
}

func NewConfig() *Config {
	return &Config{
		Client:      api.NewClient(),
		UserPokedex: make(map[string]api.Pokemon),
	}
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"map": {
			Name:        "map",
			Description: "Displays the name of  location areas in the Pokemon world",
			Callback:    Map,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the name of 20 location areas in the Pokemon world in reverse order",
			Callback:    Mapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a specific area location",
			Callback:    Explore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a Pokemon",
			Callback:    Catch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a user Pokemon",
			Callback:    Inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List of all the names of the Pokemon the user has caught",
			Callback:    Pokedex,
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
