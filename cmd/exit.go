package cmd

import (
	"fmt"
	"os"
)

func CommandExit() error {
	fmt.Println("\nExiting Pokedex . . .")
	os.Exit(1)
	return nil
}
