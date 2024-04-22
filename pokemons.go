package ecolePokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "https://pokebuildapi.fr/api/v1/pokemon/"

type PokemonResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	ApiTypes []ApiType `json:"apiTypes"`
	Stats    Stats `json:"stats"`

	// Ajoutez d'autres champs selon vos besoins
}

type Stats struct {
	HP int `json:"hp"`
	Attack int `json:"attack"`
	Defense int `json:"defense"`
	SpecialAttack int `json:"specialAttack"`
	SpecialDefense int `json:"specialDefense"`
	Speed int `json:"speed"`
}

type ApiType struct {
	Name string `json:"name"`
	Image string `json:"image"`
}

func GetAllPokemons() ([]PokemonResponse, error) {
	var pokemons []PokemonResponse
	resp, err := http.Get(url)
	if err!= nil {
        return pokemons, err
    }
	defer resp.Body.Close()
	if resp.StatusCode!= http.StatusOK {
        return pokemons, fmt.Errorf("erreur lors de la récupération de la liste des Pokémons")
    }

	if err := json.NewDecoder(resp.Body).Decode(&pokemons); err != nil {
		return pokemons, err
	}
	
	return pokemons, nil	
}

func GetPokemon(id string) (PokemonResponse, error) {
	var pokemon PokemonResponse

	response, err := http.Get(url + id)
	if err != nil {
		return pokemon, err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
		return pokemon, err
	}
	return pokemon, nil
}