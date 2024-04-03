package ecolePokemon

import (
    "encoding/json"
    "fmt"
    "io"
	"net/http"
	"strings"
)

func GetDates(pokemon Pokemon) []string {
	var dates Date
	url := pokemon.DatesLink
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error HTTP in GetDates :", err)
        return nil
	}
	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
        fmt.Println("Reading Error in GetDates :", err2)
        return nil
    }

	err3 := json.Unmarshall(body, &dates)
	if err3 != nil {
        fmt.Println("Error Unmarshall in GetDates :", err3)
        return nil
    }

	for i, date := range dates.Dates {
		if string(date[0]) == "*" { 
			dates.Dates[i] = date[1:]
		}
	}

	return dates.Dates
}

func DatesToPokemons(search string, pokemonList []Pokemon) []Pokemon {
	var pokemonsFound []Pokemon
	var isDate bool

	for _, pokemon := range pokemonList {
		isDate = false
		datesList := GetDates(pokemon)
		for _, pokemonDates := range datesList {
			if strings.Contains(strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(pokemonDates, "-", " ")"_", " ")), strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(search, "-", " "), "_", " "))) {
                isDate = true
            }
		}
		if isDate {
            pokemonsFound = append(pokemonsFound, pokemon)
        }
	}
	return pokemonsFound
}