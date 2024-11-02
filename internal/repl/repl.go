package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jamesfulreader/pokedexcli/cmd"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := cmd.GetCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)
		input = strings.TrimSpace(input)

		command, exists := commands[input]
		if !exists {
			fmt.Printf("unknown command: %s\n", input)
			fmt.Println()
			cmd.CommandHelp()
			continue
		}

		err := command.Callback()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}
	}
}
