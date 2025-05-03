package commands

import (
	"fmt"
)

func Mapb(c *Config) error {
	if c.PreviousLocationAreaUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := c.Client.GetLocationAreas(c.PreviousLocationAreaUrl)
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
