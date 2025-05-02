package commands

import (
	"fmt"

	"github.com/mgyosbel/pokedex/api"
)

func Map(c *Config) error {
	url := c.Next
	if c.Next == "" {
		url = api.LOCATION_AREA_URL
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
