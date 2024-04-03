package main

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"text/template"

	"github.com/Christianngono/ecolePokemon"
)

type Pokemon struct {
	Name           string  `json:"name"`
	ID             int     `json:"id"`
	Pokemon        string  `json:"pokemon"`
	CreationDate   int     `json:"creationDate"`
	DatesLink      string  `json:"url"`
	EvolutionsLink string  `json:"urlEvolution"`
	GenerationLink string  `json:"urlGeneration"`
	Height         float64 `json:"height"`
	Weight         float64 `json:"weight"`
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
		BackDefault  string `json:"back_default"`
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
	} `json:"moves"`
}

type Dates struct {
	// Dates located in Relation info
	DatesId   string `json:"datesid"`
	DatesName string `json:"datesname"`
}

type Date struct {
	Dates []string `json:"dates"`
}

type Location struct {
	Contry      string    `json:"contry"`
	City        string    `json:"city"`
	Locations   []string  `json:"locations"`
	MapLink     string    `json:"maplink"`
	PokemonList []Pokemon `json:"pokemonList"`
	Location    string    `json:"location"`
}

type Evolution struct {
	Evolutions   []string  `json:"evolutions"`
	PokemonList  []Pokemon `json:"pokemonList"`
	UrlEvolution string    `json:"urlEvolution"`
	Evolution    string    `json:"evolution"`
	EvolutionID  string    `json:"evolutionID"`
}

type Generation struct {
	Generations    []string  `json:"generations"`
	PokemonList    []Pokemon `json:"pokemonList"`
	UrlGenerations string    `json:"urlGenerations"`
	Genreration    string    `json:"generation"`
	GenerationID   string    `json:"generationID"`
}

type DatePage struct {
	// Info shown in the date page
	Date        string    `json:"date"`
	PokemonList []Pokemon `json:"pokemonList"`
}

type InfoPage struct {
	Location Id string
	LocationName string
	Dates []Dates
}

type Language struct {
	Languages   []string  `json:"languages"`
	PokemonList []Pokemon `json:"pokemonList"`
	LanguageID  string    `json:"language"`
	Language    string    `json:"language"`
}

const (
	url            = "https://pokeapi.co/api/v2/pokemon/ditto"
	urlEvolution   = "https://pokeapi.co/api/v2/evolution-chain/{id}/"
	urlGenerations = "https://pokeapi.co/api/v2/generation/{id or name}/"

	maplink = "https://pokeapi.co/api/v2/location/{id or name}/"
)

func main() {
	fmt.Println("Server successfully started at http://localhost:8080")
	fileServer := http.fileServer(http.Dir("static"))                 // set path to static files (images and css file)
	http.handle("/static/", http.StripPrefix("/static/", fileServer)) // connects the static files to the server
	http.HandleFunc("/index", HomeHandler)
	http.HandleFunc("/pokemon", PokemonHandler)
	http.HandleFunc("/location", LocationHandler)
	http.HandleFunc("/datepage", DatePageHandler)
	http.HandleFunc("/evolution", EvolutionHandler)
	http.HandleFunc("/generation", GenerationHandler)
	http.HandleFunc("/language", LanguageHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var frontend = template.Must(template.ParseFiles("template/index.html"))
	var pokemonList = ecolePokemon.GetAllPokemons()
	search := r.FormValue("search")
	sort := r.FormValue("sort")

	searchedPokemonList := ecolePokemon.SearchPokemons(search, pokemonList)
	if len(searchedPokemonList) == 0 {
		searchedPokemonList = ecolePokemon.TypeToPokemons(search, pokemonList)
	}

	if len(searchedPokemonList) != 0 {
		pokemonList = searchedPokemonList
	}

	if sort == "A > Z" {
		pokemonList = ecolePokemon.SortPokemon(pokemonList)
	} else if sort == "Z > A" {
		pokemonList = ecolePokemon.SortPokemon(pokemonList)
		slices.Reverse(pokemonList)
	} else if sort == "Old > New" {
		pokemonList = ecolePokemon.SortCreationDate(pokemonList)
	} else if sort == "New > Old" {
		pokemonList = ecolePokemon.SortCreationDate(pokemonList)
		slices.Reverse(pokemonList)
	}

	frontend.Execute(w, pokemonList) // send all info to  webpage
}

func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	var frontend = template.Must(template.ParseFiles("template/pokemon.html")) // connects template to html webpage
	pokemonID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	pokemon := ecolePokemon.GetPokemon(pokemonID)
	frontend.Execute(w, pokemon)
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	var frontend = template.Must(template.ParseFiles("template/location.html"))
	var pokemonList = ecolePokemon.GetAllPokemons()

	location := r.URL.Query().Get("city_country")
	pokemonList = ecolePokemon.LocationToPokemons(location, pokemonList)

	search := r.FormValue("search")
	sort := r.FormValue("sort")

	searchedPokemonList := ecolePokemon.SearchPokemons(search, pokemonList)
	if len(searchedPokemonList) == 0 {
		searchedPokemonList = ecolePokemon.LocationToPokemons(search, pokemonList)
	}

	if len(searchedPokemonList) == 0 {
		searchedPokemonList = ecolePokemon.DatesToPokemons(search, pokemonList)
	}

	if len(searchedPokemonList) != 0 {
		pokemonList = searchedPokemonList
	}

	if sort == "A > Z" {
		pokemonList = ecolePokemon.SortPokemon(pokemonList)
	} else if sort == "Z > A" {
		pokemonList = ecolePokemon.SortPokemon(pokemonList)
		slices.Reverse(pokemonList)
	} else if sort == "Old > New" {
		pokemonList = ecolePokemon.SortCreationDate(pokemonList)
	} else if sort == "New > Old" {
		pokemonList = ecolePokemon.SortCreationDate(pokemonList)
		slices.Reverse(pokemonList)
	}

	pageInfo := ecolePokemon.LocationHandler{
		Location:    ecolePokemon.FormatLocation(location),
		Maplink:     ecolePokemon.GetMapLink(location),
		PokemonList: pokemonList,
	}
	frontend.Execute(w, pageInfo)
}

func DatePageHandler(w http.ResponseWriter, r *http.Request) {
	var frontend = template.Must(template.ParseFiles("template/date.html")) // connects template to html webpage
	var pokemonList = ecolePokemon.GetAllPokemons()

	date := r.URL.Query().Get("date") // get the date's location

	pokemonList = ecolePokemon.DatesToPokemons(date, pokemonList)

	search := r.FormValue("search") // take the user input from the form element in the webpage
	sort := r.FormValue("sort")     // take the user input from the form element in the webpage

	searchedPokemonList := ecolePokemon.SearchPokemons(search, pokemonList)
	if len(searchedPokemonList) == 0 {
		searchedPokemonList = ecolePokemon.LocationsToPokemons(search, pokemonList)
	}

	if len(searchedPokemonList) == 0 {
		searchedPokemonList = ecolePokemon.DatesToPokemons(search, pokemonList)
	}

	if len(searchedPokemonList) != 0 {
		pokemonList = searchedPokemonList
	}

	if sort == "A > Z" {
		pokemonList = ecolePokemon.SortPokemon(pokemonList)
	} else if sort == "Z > A" {
		pokemonList = ecolePokemon.SortPokemon(pokemonList)
		slices.Reverse(pokemonList)
	} else if sort == "Old > New" {
		pokemonList = ecolePokemon.SortCreationDate(pokemonList)
	} else if sort == "New > Old" {
		pokemonList = ecolePokemon.SortCreationDate(pokemonList)
		slices.Reverse(pokemonList)
	}

	pageInfo := ecolePokemon.DatePage{
		Date:            ecolePokemon.FormatDate(date),
		PokemonListList: pokemonList,
	}

	frontend.Execute(w, pageInfo) // send all info to webpage
}

func EvolutionHandler(w http.ResponseWriter, r *http.Request) {
	var frontend = template.Must(template.ParseFiles("template/evolution.html"))
	evolutionID, _ := stconv.Atoi(r.URL.Query().Get("id"))
	evolution := ecolePokemon.GetEvolution(evolutionID)
	fronted.Execute(w, evolution)
}

func GenerationHandler(w http.ResponseWriter, r *http.Request) {
	var frontend = template.Must(template.ParseFile("template/generation.html"))
	generationID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	generation := ecolePokemon.GetGeneration(generationID)
	frontend.Execute(w, generation)
}

func LanguageHandler(w http.ResponseWriter, r *http.Request) {
	var fontend = template.Must(template.ParseFile("template/language.html"))
	languageID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	language := ecolePokemon.GetLanguage(languageID)
	frontend.Execute(w, language)
}
