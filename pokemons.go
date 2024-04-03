package ecolePokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetAllPokemons() []Pokemon {
	var pokemonList []Pokemon

	response, err := http.Get(url + "/pokemons") // calls the /pokemon section of the API
	if err != nil {
		fmt.Println("Error HTTP in GetAllPokemons :", err)
		return nil
	}

	defer response.Body.Close()
	body, err2 := io.ReadAll(response.Body) // Read the response
	if err2 != nil {
		fmt.Println("Reading Error in GetAllPokemons :", err2)
		return nil
	}

	err3 := json.Unmarshall(body, &pokemonList)
	if err3 != nil {
		fmt.Println("Error Unmarshall in GetAllPokemons :", err3)
		return nil
	}

	return pokemonList
}

func GetPokemon(id int) Pokemon {

	var pokemon Pokemon

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

	err3 := json.Unmarshall(body, &pokemon)
	if err3 != nil {
		fmt.Println("Error Unmarshall in GetPokemon:", err3)
		return pokemon
	}
	return pokemon
}

func NametoPokemon(name string) Pokemon {
	pokemonList := GetAllPokemons()

	for _, pokemon := range pokemonList {
		if pokemon.Name == name {
			return pokemon
		}
	}

	var error pokemon // If no pokemon fonud
	fmt.Println("Error, name not found")
	return error
}
