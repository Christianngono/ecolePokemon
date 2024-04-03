package ecolePokemon

import (
	"slices"
	"strings"
)

func sortPokemon(pokemonList []Pokemon) []Pokemon {
	var sortedPokemonList []Pokemon

	for _, pokemon := range pokemonList {
		var isLast = true
		if len(sortedPokemonList) == 0 {
			sortedPokemonList = append(sortedPokemonList, pokemon)
		} else {
			for i, sortedPokemon := range sortedPokemonList {
				if pokemon.Name <= sortedPokemon.Name {
					sortedPokemonList = slices.Insert(sortedPokemonList, i, pokemon)
					isLast = false
					break
				}
			}
			if isLast { 
                sortedPokemonList = append(sortedPokemonList, pokemon)
            }
		}
	}
	return sortedPokemonList
}

func sortCreationDate(pokemonList []Pokemon) []Pokemon {
	var sortedPokemonList []Pokemon

    for _, pokemon := range pokemonList {
        var isLast = true
        if len(sortedPokemonList) == 0 {
            sortedPokemonList = append(sortedPokemonList, pokemon)
        } else {
            for i, sortedPokemon := range sortedPokemonList {
                if pokemon.CreationDate <= sortedPokemon.CreationDate {
                    sortedPokemonList = slices.Insert(sortedPokemonList, i, pokemon)
                    isLast = false
                    break
                }
            }
            if isLast { 
                sortedPokemonList = append(sortedPokemonList, pokemon)
            }
        }
    }
    return sortedPokemonList
}

func SearchPokemons(search string, pokemonList []Pokemon) []Pokemon {
	var searchedPokemonList []Pokemon

    for _, pokemon := range pokemonList {
        if strings.Contains(strings.ToLower(pokemon.Name), strings.ToLower(strings.TrimSpace(input))) {
            searchedPokemonList = append(searchedPokemonList, pokemon)
        }
    }
    return searchedPokemonList
}