package ecolePokemon

import (
	"encoding/json"
    "fmt"
    "io"
    "net/http"
    "slices"
    "strings"
)

func GetLanguages(pokemon Pokemon) []string {
	var languages Language
	url := pokemon.LanguagesLink
	response, err := http.Get(url)
	if err != nil {
        fmt.Println("Error HTTP in GetLanguages:", err)
        return nil
    }
	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
        fmt.Println("Reading Error in GetLanguages:", err2)
        return nil
    }
	err3 := json.Unmarshall(body, &languages)
	if err3 != nil {
        fmt.Println("Error Unmarshall in GetLanguages:", err3)
        return nil
    }
	return languages.Languages
}

func LanguagesToPokemons(languages []string) []Pokemon {
	var sortedPokemon []Pokemon
	var isIn bool

	for _, language := range languages {
		isIn = false
		languagesList := GetLanguages(pokemon) {
			for _, languagePokemon := range languagesList {
                if strings.Contains(strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(languagePokemon, "-", " "), "_", " ")), strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(language, "-", " "), "_", " "))) {
                    isIn = true
                }
            }
            if isIn {
                sortedPokemonList = append(sortedPokemonList, languagePokemon)
            }
		}
		return sortedPokemonList 
	}
}

func GetLanguagesLink(pokemon Pokemon) string {
	Languages := []string{"Afrikaan", "English", "French", "German", "Italian", "Spanish", "Chinese", "Japanese", "Korean", "Polish", "Portuguese", "Russian", "Swedish", "Turkish", "Vietnamese", "Indonesian", "Arabic"}
	if slices.Contains(Languages, language) {
		return urlLanguages1 + language + urlLanguages2 + pokemon + urlLanguages3
    } else {
		return urlLanguages1 + language + urlLanguages2 + pokemon + urlLanguages3
	}
}

