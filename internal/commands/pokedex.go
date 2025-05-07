package commands

import "fmt"

func Pokedex(c *Config) error {
	fmt.Println("Your Pokedex:")
	for _, p := range c.UserPokedex {
		fmt.Printf("  - %s\n", p.Name)
	}
	return nil
}
