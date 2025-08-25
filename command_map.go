package main

import (
	"fmt"
)

func mapCallback(cfg *config) error {
	respLocationArea, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = respLocationArea.Next
	cfg.prevLocationsURL = respLocationArea.Previous

	fmt.Println()
	for _, locationArea := range respLocationArea.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println()

	return nil
}

func mapbCallback(cfg *config) error {
	respLocationArea, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = respLocationArea.Next
	cfg.prevLocationsURL = respLocationArea.Previous

	fmt.Println()
	for _, locationArea := range respLocationArea.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println()

	return nil
}
