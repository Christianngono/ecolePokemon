package ecolePokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetVersions(id int) string {
	var version string

	response, err := http.Get(url + "/versions/" + strconv.Itoa(id))
	if err != nil {
		fmt.Println("Error HTTP in GetVersions:", err)
		return version
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Reading Error in GetVersions:", err)
		return version
	}

	err = json.Unmarshal(body, &version)
	if err != nil {
		fmt.Println("Error Unmarshal in GetVersions:", err)
		return version
	}

	return version
}
