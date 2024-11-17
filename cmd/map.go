package cmd

import (
	"fmt"
)

func CommandMap([]string) error {
	fmt.Println()

	var pageURL *string
	if config.NextURL != "" {
		pageURL = &config.NextURL
	}

	locationURL, err := config.Client.GetLocations(pageURL)
	if err != nil {
		return err
	}

	if locationURL.Next != nil {
		config.NextURL = *locationURL.Next
	}
	if locationURL.Previous != nil {
		config.PrevURL = *locationURL.Previous
	}

	for _, locationItem := range locationURL.Results {
		fmt.Println(locationItem.Name)
	}

	fmt.Println()
	return nil
}

func CommandMapB([]string) error {
	fmt.Println()
	if config.PrevURL == "" {
		fmt.Println()
		fmt.Println("cannot go back any further")
		fmt.Println()
		return nil
	}

	var pageURL *string
	if config.PrevURL != "" {
		pageURL = &config.PrevURL
	}

	locationURL, err := config.Client.GetLocations(pageURL)
	if err != nil {
		return err
	}

	if locationURL.Next != nil {
		config.NextURL = *locationURL.Next
	}
	if locationURL.Previous != nil {
		config.PrevURL = *locationURL.Previous
	}

	for _, locationItem := range locationURL.Results {
		fmt.Println(locationItem.Name)
	}
	fmt.Println()
	return nil
}
