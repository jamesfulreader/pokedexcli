package cmd

import "fmt"

func CommandExplore(location []string) error {
	fmt.Println("I wil explore: ", location)
	return nil
}
