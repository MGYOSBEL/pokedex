package commands

import "fmt"

func Help(*Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("")
	for _, c := range GetCommands() {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}
	return nil
}
