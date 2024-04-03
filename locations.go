package ecolePokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"stings"
)


func GetLocations(pokemon Pokemon) []string {
	var locations []string
	// get the location of the pokemon
	url := pokemon.MapLink
	// Calls the location section of the pokeapi
	response, err := http.Get(url)
	if err!= nil {
        fmt.Println("Error HTTP in GetLocations:", err)
        return nil
    }
	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body)
	if err2!= nil {
        fmt.Println("Reading Error in GetLocations:", err2)
        return nil
    }
	err3 := json.Unmarshall(body, &pokemon)
	if err3 != nil {
        fmt.Println("Error Unmarshall in GetLocations:", err3)
        return nil
    }
	return pokemon.Locations	
}

func LocationToPokemons(location string, pokemonList []Pokemon) []Pokemon {
	var sortedPokemonList []Pokemon
	var isIn bool

	for _, pokemon := range pokemonList {
		isIn = false
		locationsList := GetLocations(pokemon)
		for _, location := range locationsList {
            if strings.Contains(strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(pokemonLocation, "-", " ")"_", " ")), strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(location, "-", " "), "_", " "))) {
                isIn = true
            }
        }
		if isIn {
			sortedPokemonList = append(sortedPokemonList, pokemon)
		}
	}
	return sortedPokemonList
}

func GetMapLink(location string) string {
	States := []string{"Alabama", "Alaska", "Arizona", " California", " Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", " Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan"}
	if slices.Contains(States, location) {
		return maplink1 + location + maplink2 + "6" + maplink3
	} else {
		return maplink1 + location + maplink2 + "10" + maplink3
	}

}