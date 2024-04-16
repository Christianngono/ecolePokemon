package ecolePokemon

import (
	"slices"
	"strings"
)

func SortPokemons(pokemonList []string) []string {
	var sortedPokemonList []string

	for _, pokemon := range pokemonList {
		var isLast = true
		if len(sortedPokemonList) == 0 {
			sortedPokemonList = append(sortedPokemonList, pokemon)
		} else {
			for i, sortedPokemon := range sortedPokemonList {
				if pokemon <= sortedPokemon {
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

func SortVersions(pokemonForm []string) []string {
	var sortedPokemonList []string

	for _, pokemon := range pokemonForm {
		var isLast = true
		if len(sortedPokemonList) == 0 {
			sortedPokemonList = append(sortedPokemonList, pokemon)
		} else {
			for i, sortedPokemon := range sortedPokemonList {
				if pokemon <= sortedPokemon {
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

func SearchPokemons(input string, pokemonList []string) []string {
	var searchedPokemonList []string

	for _, pokemon := range pokemonList {
		if strings.Contains(strings.ToLower(pokemon), strings.ToLower(strings.TrimSpace(input))) {
			searchedPokemonList = append(searchedPokemonList, pokemon)
		}
	}
	return searchedPokemonList
}
