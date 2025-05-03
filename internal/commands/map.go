package commands

import (
	"fmt"
)

func Map(c *Config) error {
	res, err := c.Client.GetLocationAreas(c.NextLocationAreaUrl)
	if err != nil {
		return err
	}

	c.NextLocationAreaUrl = res.Next
	c.PreviousLocationAreaUrl = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
