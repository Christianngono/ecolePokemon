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
	ResistanceModifyingAbilitiesForApis []ResistanceModifyingAbilitiesForApi `json:"resistanceModifyingAbilitiesForApis"`
	ApiGenerations []ApiGeneration `json:"apiGenerations"`
	ApiResistances []ApiResistance `json:"apiResistances"`	
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

type ResistanceModifyingAbilitiesForApi struct {
    Name string `json:"name"`
    Image string `json:"image"`
}
type ApiGeneration struct {
    Name string `json:"name"`
    Image string `json:"image"`
}


type ApiResistance struct {
	Name string `json:"name"`
	Image string `json:"image"`
	DamageMultiplier float64 `json:"damage_multipliers"`
	DamageRelation string `json:"damage_relation"`
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

func GetPokemonByName(name string) (PokemonResponse, error) {
	var pokemon PokemonResponse

    response, err := http.Get(url + name)
    if err != nil {
        return pokemon, err
    }
    defer response.Body.Close()

    if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
        return pokemon, err
    }
    return pokemon, nil
}

func GetAllPokemonsStats()([]PokemonResponse, error) {
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

func GetPokemonStatsByName(name string) (PokemonResponse, error) {
	var pokemon PokemonResponse

    response, err := http.Get(url + name)
    if err != nil {
        return pokemon, err
    }
    defer response.Body.Close()

    if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
        return pokemon, err
    }
    return pokemon, nil
}

func GetAllPokemonsResistanceModifyingAbilitiesForApis()([]PokemonResponse, error) {
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

func GetPokemonResistanceModifyingAbilitiesForApisByName(name string) (PokemonResponse, error) {
	var pokemon PokemonResponse

    response, err := http.Get(url + name)
    if err != nil {
        return pokemon, err
    }
    defer response.Body.Close()

    if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
        return pokemon, err
    }
    return pokemon, nil
}

func GetAllPokemonsApiGenerations()([]PokemonResponse, error) {
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

func GetPokemonApiGenerationsByName(name string) (PokemonResponse, error) {
	var pokemon PokemonResponse

    response, err := http.Get(url + name)
    if err != nil {
        return pokemon, err
    }
    defer response.Body.Close()

    if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
        return pokemon, err
    }
    return pokemon, nil
}

func GetAllPokemonsApiResistances()([]PokemonResponse, error) {
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

func GetPokemonApiResistancesByName(name string) (PokemonResponse, error) {
	var pokemon PokemonResponse

    response, err := http.Get(url + name)
    if err != nil {
        return pokemon, err
    }
    defer response.Body.Close()

    if err := json.NewDecoder(response.Body).Decode(&pokemon); err != nil {
        return pokemon, err
    }
    return pokemon, nil
}