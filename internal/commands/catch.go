package commands

import (
	"fmt"
	"math/rand"
)

const BASE_DIFFICULTY = 10

func Catch(c *Config) error {
	name := c.CommandArguments[0]
	if _, ok := c.UserPokedex[name]; ok {
		fmt.Printf("%s is already in your pocket\n", name)
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := c.Client.GetPokemonByName(&name)
	if err != nil {
		return err
	}

	if caught := randomize(pokemon.BaseExperience); !caught {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	c.UserPokedex[name] = pokemon
	fmt.Printf("%s was caught!\n", name)
	return nil
}

func randomize(exp int) bool {
	return rand.Intn(exp%BASE_DIFFICULTY) == 0
}
