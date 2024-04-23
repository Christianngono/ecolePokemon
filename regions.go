package ecolePokemon

import (
	"encoding/json"
	"fmt"
	"net/http"

)

// RegionToPokemons trouve les pokemons associés à une région donnée
func RegionToPokemons(region string, Pokemons []string) map[string][]string {

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

func GetRegions()([]string, error) {
	var regions []string
	resp, err := http.Get(urlRegion)
	if err!= nil {
        return regions, err
    }
	defer resp.Body.Close()
	if resp.StatusCode!= http.StatusOK {
        return regions, fmt.Errorf("erreur lors de la récupération de la liste des régions")
    }
	if err := json.NewDecoder(resp.Body).Decode(&regions); err != nil {
        return regions, err
    }
	return regions, nil	
}

func GetRegion(region string) (string, error) {
	var regionName string

    response, err := http.Get(urlRegion + region)
    if err != nil {
        return regionName, err
    }
    defer response.Body.Close()

    if err := json.NewDecoder(response.Body).Decode(&regionName); err != nil {
        return regionName, err
    }
    return regionName, nil
}




