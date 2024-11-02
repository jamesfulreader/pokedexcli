package cmd

import "fmt"

func CommandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := GetCommands()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()

	return nil
}
