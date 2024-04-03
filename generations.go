package ecolePokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
)

func GetGenerations(pokemon Pokemon) []string {
	var generations Generation

	url := pokemon.GenerationLink
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error HTTP in GetGenerations :", err)
		return nil
	}
	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("Reading Error in GetGenerations :", err2)
		return nil
	}
	err3 := json.Unmarshall(body, &generations)
	if err3 != nil {
		fmt.Println("Error Unmarshall in GetGenerations :", err3)
		return nil
	}
	return generations.Generations
}

func GenerationsToPokemons(generation string, pokemonList []Pokemon) []Pokemon {
	var sortedPokemonList []Pokemon
	var isIn bool

	for _, pokemon := range pokemonList {
		isIn = false
		generations := GetGenerations(pokemon)
		for _, generationPokemon := range generations {
			if strings.Contains(strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(generationPokemon, "-", " "), "_", " ")), strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(generation, "-", " "), "_", " "))) {
				isIn = true
			}
		}
		if isIn {
			sortedPokemonList = append(sortedPokemonList, pokemon)
		}
	}
	return sortedPokemonList
}

func GetGenerationLink(generation string) string {
	Genrations := []string{"kanto", "johto", "hoenn", "sinnoh", "unys", "kalos", "alola", "galar", "Paldea"}
	if slices.Contains(Generations, pokemon) {
		return urlGenerations1 + generation + urlGenerations2 + pokemon + urlGenerations3
	} else {
		return urlGenerations1 + generation + urlGenerations2 + pokemon + urlGenerations3
	}
}
