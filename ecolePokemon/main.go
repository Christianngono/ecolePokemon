package main

import (
	"fmt"
	"net/http"
	"text/template"

	"ecolePokemon"
)

// Pokemon représente la structure d'un Pokémon
type Pokemon struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Region    string `json:" region"`
	Sprites   string `json:"sprites"`
	Url       string `json:"url"`
	UrlRegion string `json:"urlregion"`
}

// Region représente la structure d'une région
type Region struct {
	Name      string   `json:"name"`
	Regions   []string `json:"regions"`
	UrlRegion string   `json:"url_region"`
	Pokemons  []string `json:"pokemons"`
}

const (
	UrlRegion = "https://pokeapi.co/api/v2/region/"
)

func main() {
	fmt.Println("Server successfully started at http://localhost:8080")
	fileServer := http.FileServer(http.Dir("static")) // Utilisation d'un chemin relatif pour les fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	http.HandleFunc("/index", HomeHandler)
	http.HandleFunc("/pokemons", PokemonHandler)
	http.HandleFunc("/regions", RegionsHandler)
	http.HandleFunc("/region", RegionHandler)

	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pokemons, err := ecolePokemon.GetAllPokemons()

	if err != nil {
		// Gérer l'erreur ici, comme afficher un message d'erreur ou rediriger l'utilisateur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("template/index.html")
	if err != nil {
		// Gérer l'erreur ici, comme afficher un message d'erreur ou rediriger l'utilisateur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, pokemons)
}

func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	var frontend = template.Must(template.ParseFiles("template/pokemons.html"))
	// Récupère l'id Pokémon à partir des paramètres de la requête
	pokemonID := r.FormValue("id")

	// Vérifier si l'ID est vide
	if pokemonID == "" {
		// Récupérer l'URL de la page index
		indexURL := "/index"

		// Rediriger l'utilisateur vers la page index en cas d'erreur
		http.Redirect(w, r, indexURL, http.StatusSeeOther)
		return
	}

	// Obtenir le Pokémon avec l'ID spécifié
	pokemon, err := ecolePokemon.GetPokemon(pokemonID)
	if err != nil {
		// Gérer l'erreur ici, comme afficher un message d'erreur ou rediriger l'utilisateur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécuter le template avec les données du Pokémon
	frontend.Execute(w, pokemon)
}

func RegionsHandler(w http.ResponseWriter, r *http.Request) {
	regions, err := ecolePokemon.GetRegions()
	if err != nil {
		// Gérer l'erreur ici, comme afficher un message d'erreur ou rediriger l'utilisateur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := template.ParseFiles("template/regions.html")
	if err != nil {
		// Gérer l'erreur ici, comme afficher un message d'erreur ou rediriger l'utilisateur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, regions)
}

func RegionHandler(w http.ResponseWriter, r *http.Request) {
	var frontend = template.Must(template.ParseFiles("template/region.html"))
	// Récupère l'id région à partir des paramètres de la requête
	regionID := r.FormValue("id")

	// Vérifier si l'ID est vide
	if regionID == "" {
		// Récupérer l'URL de la page index
		indexURL := "/index"

		// Rediriger l'utilisateur vers la page index en cas d'erreur
		http.Redirect(w, r, indexURL, http.StatusSeeOther)
		return
	}

	// Obtenir la région avec l'ID du pokémon spécifié
	region, err := ecolePokemon.GetRegion(regionID)
	if err != nil {
		// Gérer l'erreur ici, comme afficher un message d'erreur ou rediriger l'utilisateur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécuter le template avec les données de la région
	frontend.Execute(w, region)

}
