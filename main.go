package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt = "Pokedex > "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		text := cleanInput(scanner.Text())
		if len(text) != 0 {
			fmt.Printf("Your command was: %s\n", text[0])
		}
		fmt.Print(prompt)
	}
}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
}
