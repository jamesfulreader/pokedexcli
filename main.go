package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "help" {
			commandHelp()
		} else if input == "exit" {
			commandExit()
		} else {
			fmt.Println("unrecognized command")
			commandHelp()
		}
	}

}
