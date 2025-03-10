package pokeapi

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	pokemoncli "pokemon-cli/internal"
	"strconv"
)

const (
	pokemonsEndpoint = "/pokemon"
	pokeapiURL       = "https://pokeapi.co/api/v2"
)

var apiResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type pokemonRepo struct {
	url string
}

func NewPokeapiRepository() pokemoncli.PokemonRepo {
	return &pokemonRepo{
		url: pokeapiURL,
	}
}

func (b *pokemonRepo) GetPokemons() (pokemons []pokemoncli.Pokemon, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, pokemonsEndpoint))
	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}

	pokemons = make([]pokemoncli.Pokemon, len(apiResponse.Results))
	for i, result := range apiResponse.Results {
		pokemons[i] = pokemoncli.Pokemon{
			Name: result.Name,
			Img:  result.URL,
			Id:   i,
		}
	}
	return
}

func (b *pokemonRepo) DumpPokemonsToCSV(pokemons []pokemoncli.Pokemon, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating CSV file: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"ID", "Name", "URL"}
	if err := writer.Write(header); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing CSV header: %v\n", err)
		return
	}

	// Write data rows
	for _, pokemon := range pokemons {
		record := []string{
			strconv.Itoa(pokemon.Id),
			pokemon.Name,
			pokemon.Img,
		}
		if err := writer.Write(record); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing pokemon to CSV: %v\n", err)
			return
		}
	}

	fmt.Printf("Successfully exported %d pokemons to %s\n", len(pokemons), fileName)
}
