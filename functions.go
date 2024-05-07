package ecolePokemon

import (
	"strings"
	"sort"
)

func SortAllPokemons(pokemons []string) []string {
	var sortedPokemons []string

	for _, pokemon := range pokemons {
		var isLast = true
		if len(sortedPokemons) == 0 {
			sortedPokemons = append(sortedPokemons, pokemon)
		} else {
			for i, sortedPokemon := range sortedPokemons {
				if pokemon <= sortedPokemon {
					sortedPokemons = Insert(sortedPokemons, i, pokemon)
					isLast = false
					break
				}
			}
			if isLast {
				sortedPokemons = append(sortedPokemons, pokemon)
			}
		}
	}
	return sortedPokemons
}

func SortPokemon(Pokemon string) string {
	var sortedPokemon string
	for _, letter := range Pokemon {
		sortedPokemon = sortedPokemon + string(letter)
		sortedPokemon = strings.ToLower(sortedPokemon)
		sortedPokemon = SortAllPokemons(strings.Split(sortedPokemon, ""))[0]
		sortedPokemon = strings.ToUpper(sortedPokemon)
	}
	return sortedPokemon
}

func SortPokemonByName(sortedPokemon []string) []string {
	var sortedPokemons []string

    for _, pokemon := range sortedPokemon {
        var isLast = true
        if len(sortedPokemons) == 0 {
            sortedPokemons = append(sortedPokemons, pokemon)
        } else {
            for i, sortedPokemon := range sortedPokemons {
                if pokemon <= sortedPokemon {
                    sortedPokemons = Insert(sortedPokemons, i, pokemon)
                    isLast = false
                    break
                }
            }
            if isLast {
                sortedPokemons = append(sortedPokemons, pokemon)
            }
        }
    }
    return sortedPokemons
}
func SortRegions(regionsMap map[string]string) map[string]string {
	// Créer une nouvelle carte pour stocker les regions triées
	var sortedRegionsMap = make(map[string]string)

	// Collecter les noms des régions
	var regions []string

	for region := range regionsMap {
		regions = append(regions, region)
	}

	// Trier les noms des régions
	 sort.Strings(regions)

	// Insérer les régions triées dans la map
	for _, region := range regions {
		sortedRegionsMap[region] = regionsMap[region]
	}

	return sortedRegionsMap
}

func SortRegion(id string) string {
	var sortedRegion string
    for _, letter := range id {
        sortedRegion = sortedRegion + string(letter)
        sortedRegion = strings.ToLower(sortedRegion)
        sortedRegion = SortAllPokemons(strings.Split(sortedRegion, ""))[0]
        sortedRegion = strings.ToUpper(sortedRegion)
    }
    return sortedRegion   
}

	
func SearchPokemons(input string, pokemons []string) []string {
	var searchedPokemons []string

	for _, pokemon := range pokemons {
		if strings.Contains(strings.ToLower(pokemon), strings.ToLower(strings.TrimSpace(input))) {
			searchedPokemons = append(searchedPokemons, pokemon)
		}
	}
	return searchedPokemons
}

func SortStrings(strs []string) []string {
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if strs[i] > strs[j] {
				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}
	return strs

}

// Insert insère une chaîne de caractères dans une slice à un index donné.
func Insert(slice []string, index int, str string) []string {
	slice = append(slice[:index+1], append([]string{str}, slice[index:]...)...)
	return slice
}

// Reverse renverse une slice de chaînes de caractères.
func Reverse(slice []string) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}
