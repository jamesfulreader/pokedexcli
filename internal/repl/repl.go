package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jamesfulreader/pokedexcli/cmd"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if input == "help" {
			cmd.CommandHelp()
			fmt.Println("Help")
		} else if input == "exit" {
			cmd.CommandExit()
		} else {
			fmt.Println("unrecognized command")
			cmd.CommandHelp()
		}
	}
}
