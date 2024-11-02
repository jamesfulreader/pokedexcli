package repl

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if input == "help" {
			commandHelp()
			fmt.Println("Help")
		} else if input == "exit" {
			commandExit()
		} else {
			fmt.Println("unrecognized command")
			commandHelp()
		}
	}
}

func commandHelp() {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:\nhelp: Displays a help message\nexit: Exit the Pokedex")
	fmt.Println()
}

func commandExit() {
	fmt.Println("Exiting . . .")
	os.Exit(1)
}
