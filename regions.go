package ecolePokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"	
)

// RegionToPokemons trouve les pokemons associés à une région donnée
func RegionToPokemons(region string, pokemonForm []string) map[string][]string {

	// Carte de correspondance entre les régions et les pokemons
	regionsMap := map[string][]string{
		"Kanto":     {"Bulbasaur", "Charmander", "Squirtle"},
		"Johto":     {"Chikorita", "Cyndaquil", "Totodile"},
		"Hoenn":     {"Treecko", "Torchic", "Mudkip"},
		"Sinnoh":    {"Turtwig", "Chimchar", "Piplup"},
		"Unova":     {"Snivy", "Tepig", "Oshawott"},
		"Kalos":     {"Chespin", "Fennekin", "Froakie"},
		"Alola":     {"Rowlet", "Litten", "Popplio"},
		"Galar":     {"Grookey", "Scorbunny", "Sobble"},
		"Isshu":     {"Victini", "Tsutarja", "Pokabu"},
		"Fiore":     {"Pikachu", "Vulpix", "Meowth"},
		"Almia":     {"Grovyle", "Chimchar", "Piplup"},
		"Oblivia":   {"Shaymin", "Deoxys", "Manaphy"},
		"Orre":      {"Espeon", "Umbreon", "Leafeon"},
		"Kanlara":   {"Eevee", "Pidgey", "Rattata"},
		"Holon":     {"Magikarp", "Feebas", "Wailmer"},
		"Fizzytopia": {"Mew", "Celebi", "Jirachi"},
	}

	// Recherche des pokemons associés à la région spécifiée
	pokemons, ok := regionsMap[region]
	if !ok {
		fmt.Println("error: region non valide")
		return nil
	}
	// Créer une nouvelle map pour les pokemons associés à la région spécifiée
	regionPokemonMap := make(map[string][]string)
    regionPokemonMap[region] = pokemons

    return regionPokemonMap

}

const urlRegion = "https://pokeapi.co/api/v2/region/"

func GetRegions(id int)string {
	var region string

	response, err := http.Get(urlRegion + strconv.Itoa(id))
	if err != nil {
		fmt.Println("Erreur HTTP lors de la récupération de la région :", err)
		return region
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur de lecture lors de la récupération de la région :", err)
		return region
	}

	err = json.Unmarshal(body, &region)
	if err != nil {
		fmt.Println("Erreur de désérialisation JSON lors de la récupération de la région :", err)
		return region
	}

	return region	
}
