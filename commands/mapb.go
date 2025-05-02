package commands

import (
	"fmt"

	"github.com/mgyosbel/pokedex/api"
)

func Mapb(c *Config) error {
	url := c.Previous
	if c.Previous == "" {
		fmt.Println("you're on the first page")
	}
	res, err := api.GetLocationAreas(url)
	if err != nil {
		return err
	}

	c.Next = res.Next
	c.Previous = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
