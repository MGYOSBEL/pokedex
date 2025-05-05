package commands

import "fmt"

func Explore(c *Config) error {
	res, err := c.Client.GetLocationAreaByName(c.CurrentLocationName)
	if err != nil {
		return err
	}
	for _, pe := range res.PokemonEncounters {
		fmt.Println(pe.Pokemon.Name)
	}
	return nil
}
