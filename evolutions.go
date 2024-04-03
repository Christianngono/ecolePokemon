package ecolePokemon

import (
	"encoding/json"
    "fmt"
    "io"
    "net/http"
	"slices"
    "strings"
)

func GetEvolutions(pokemon Pokemon) []string {
	var evolutions evolution 
	url := pokemon.EvolutionsLink
	response, err := http.Get(url)
	if err != nil {
        fmt.Println("Error HTTP in GetEvolutions :", err)
        return nil
    }
	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
        fmt.Println("Reading Error in GetEvolutions :", err2)
        return nil
    }
	err3 := json.Unmarshall(body, &evolutions)
	if err3 != nil {
        fmt.Println("Error Unmarshall in GetEvolutions :", err3)
        return nil
    }
	return evolutions.Evolutions
}

func EvolutionsToPokemons(evolution string, pokemonList []Pokemon) []Pokemon {
	var sortedPokemonList []Pokemon
    var isIn bool
    for _, pokemon := range pokemonList {
		isIn = false
		evolutions := GetEvolutions(pokemon)
        for _, evolutionPokemon := range evolutions {
			if strings.Contains(strings.Tolower(strings.ReplaceAll(strings.ReplaceAll(evolutionPokemon, "-", " "), "_", " ")), strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(evolution, "-", " "), "_", " "))) {
				isIn = true
			}    
        }
		if isIn {
            sortedPokemonList = append(sortedPokemonList, pokemon)
        }
    }
    return sortedPokemonList
}

Get EvolutionLink(evolution string) string {
	var url string
    if slices.Contains(evolutions, evolution) {
		return urlEvolution1 + evolution + urlEvolution2 + "6" urlEvolution3	
	} else {
		return urlEvolution1 + evolution + urlEvolution2 + "10" urlEvolution3
	}
}