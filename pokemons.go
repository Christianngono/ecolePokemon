package ecolePokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const url = "https://pokeapi.co/api/v2/pokemon/ditto"

func GetAllPokemons() []string {
	var pokemonList []string


	response, err := http.Get(url + "/pokemons") // calls the /pokemon section of the API
	if err != nil {
		fmt.Println("Error HTTP in GetAllPokemons :", err)
		return nil
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body) // Read the response
	if err != nil {
		fmt.Println("Reading Error in GetAllPokemons :", err)
		return nil
	}

	err3 := json.Unmarshal(body, &pokemonList)
	if err3 != nil {
		fmt.Println("Error Unmarshal in GetAllPokemons :", err3)
		return nil
	}

	return pokemonList
}

func GetPokemon(id int) string {

	var pokemon string

	response, err := http.Get(url + "/pokemons/" + strconv.Itoa(id))
	if err != nil {
		fmt.Println("Error HTTP in GetPokemon:", err)
		return pokemon
	}

	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("Reading Error in GetPokemon:", err2)
		return pokemon
	}

	err3 := json.Unmarshal(body, &pokemon)
	if err3 != nil {
		fmt.Println("Error Unmarshal in GetPokemon:", err3)
		return pokemon
	}
	return pokemon
}

func NametoPokemon(name string) string {
	pokemonList := GetAllPokemons()

	for _, pokemon := range pokemonList {
		if pokemon == name {
			return pokemon
		}
	}

	var error string // If no pokemon fonud
	fmt.Println("Error, name not found")
	return error
}
