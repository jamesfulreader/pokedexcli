package cmd

import (
	"fmt"
	"os"
)

func CommandExit([]string) error {
	fmt.Println("\nExiting Pokedex . . .")
	os.Exit(1)
	return nil
}
