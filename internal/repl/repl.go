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

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		commandName := strings.ToLower(args[0])

		command, exists := commands[commandName]
		if !exists {
			fmt.Printf("unknown command: %s\n", input)
			fmt.Println()
			cmd.CommandHelp([]string{})
			continue
		}

		err := command.Callback(args)
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
			fmt.Println("enter \"help\" for available commands")
		}
	}
}
