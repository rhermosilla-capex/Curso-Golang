package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	pokemon "pokemon-concurrency/internal"
)

const (
	pokeapiURL            = "https://pokeapi.co/api/v2"
	pokemonsEndpoint      = "/pokemon"
	pokemonDetailEndpoint = "/pokemon/%v"
)

var pokemonsApiResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var pokemonsDetailApiResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
}

type pokemonRepo struct {
	url string
}

func NewPokeapiRepository() pokemon.PokemonRepo {
	return &pokemonRepo{
		url: pokeapiURL,
	}
}

func (b *pokemonRepo) GetPokemons() (pokemons []pokemon.Pokemon, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, pokemonsEndpoint))
	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &pokemonsApiResponse)
	if err != nil {
		return nil, err
	}

	pokemons = make([]pokemon.Pokemon, len(pokemonsApiResponse.Results))
	for i, result := range pokemonsApiResponse.Results {
		pokemons[i] = pokemon.NewPokemon(i, result.Name)
	}
	return
}

func (b *pokemonRepo) GetPokemonDetail(id int) (pok pokemon.PokemonDetail, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, fmt.Sprintf(pokemonDetailEndpoint, id)))
	if err != nil {
		return pokemon.PokemonDetail{}, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return pokemon.PokemonDetail{}, err
	}

	err = json.Unmarshal(contents, &pokemonsDetailApiResponse)
	if err != nil {
		return pokemon.PokemonDetail{}, err
	}

	abilities := make(pokemon.Abilities, len(pokemonsDetailApiResponse.Abilities))
	for i, ability := range pokemonsDetailApiResponse.Abilities {
		abilityInfo := ability.Ability
		abilities[i] = pokemon.NewAbilityDetail(
			ability.IsHidden,
			ability.Slot,
			pokemon.NewAbility(abilityInfo.Name, abilityInfo.URL),
		)
	}

	pok = pokemon.PokemonDetail{
		Id:        pokemonsDetailApiResponse.Id,
		Name:      pokemonsDetailApiResponse.Name,
		Height:    pokemonsDetailApiResponse.Height,
		Weight:    pokemonsDetailApiResponse.Weight,
		Abilities: abilities,
	}
	return
}
