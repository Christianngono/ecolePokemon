package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Pokemon struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Height         float64 `json:"height(m)"`
	Weight         float64 `json:"weight(kg)"`
	BaseExperience int     `json:"base_experience"`
	Abilities      []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
		// Add more image fields or sprites

	} `json:"sprites"`
	Stats []struct {
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
	Moves struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
		// Add more movement fields here if needed (e.g., "version_group_details")

	} `json:"moves"`
	// Add other fields like capacities, types, and evolutions as necessary
}

func main() {

	r := mux.NewRouter()
	// Pages
	r.HandleFunc("/index", HomeHandler)
	r.HandleFunc("/pokemon/{name}", pokemonHandler)
	r.HandleFunc("/api/pokemon{name}", apiPokemonHandler)

	http.Handle("/", r)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome on ecolePokemon!")
}

func pokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonName := vars["name"]
	pokemon, err := fetchPokemonDetails(pokemonName)
	if err != nil {
		errorHandler(w, "Error fetching Pokemon Details")
		return
	}
	renderTemplate(w, "pokemonDetails.html", pokemon)
}

func apiPokemonHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonName := vars["name"]
	pokemon, err := fetchPokemonDetails(pokemonName)
	if err != nil {
		http.Error(w, "Failed retrieving Pokémon list", 400)
		return
	}
	json.NewEncoder(w).Encode(pokemon)
}
func fetchPokemonDetails(name string) (*Pokemon, error) {
	BaseURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	res, err := http.Get(BaseURL)
	if err != nil || res.StatusCode != 200 {
		return nil, errors.New("failed getting pokemon details")
	}

	defer res.Body.Close()
	var pokemon Pokemon
	if err := json.NewDecoder(res.Body).Decode(&pokemon); err != nil {
		return nil, errors.New("error decoding JSON")
	}
	return &pokemon, nil
}

// Error handler for the application
func errorHandler(w http.ResponseWriter, message string) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, map[string]interface{}{"Message": message})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Error loading template.", http.StatusInternalServerError)
		errorHandler(w, "Error loading template.")
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute template.", http.StatusInternalServerError)
		errorHandler(w, "Failed to execute template.")
		return
	}
}
