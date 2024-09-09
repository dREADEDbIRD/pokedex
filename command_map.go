package main

import (
    "errors"
    "fmt"
)


func commandMapf(cfg *config) error {
    locationsResp, err := cfg.pokeapitClient.ListLocations(cfg.nextLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = locationsResp.Next
    cfg.prevLocationsURL = locationResp.Previous

    for _, loc := range locationsResp.Results {
        fmt.Println(loc.Name)
    }
    return nil
}

func commandMapb
