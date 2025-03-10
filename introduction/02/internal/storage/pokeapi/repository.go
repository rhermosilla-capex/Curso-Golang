package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	pokemoncli "pokemon-cli/internal"
)

const (
	pokemonsEndpoint = "/pokemon"
	pokeapiURL       = "https://pokeapi.co/api/v2"
)

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

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &pokemons)
	if err != nil {
		return nil, err
	}
	return
}
