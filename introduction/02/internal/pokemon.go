package pokemoncli

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"url"`
}

type PokemonRepo interface {
	GetPokemons() ([]Pokemon, error)
	DumpPokemonsToCSV(pokemons []Pokemon, fileName string)
}

func NewPokemon(id int, name, img string) Pokemon {
	return Pokemon{
		Id:   id,
		Name: name,
		Img:  img,
	}
}
