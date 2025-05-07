package commands

import "fmt"

func Inspect(c *Config) error {
	name := c.CommandArguments[0]
	pokemon, ok := c.UserPokedex[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("%s", pokemon)
	return nil
}
