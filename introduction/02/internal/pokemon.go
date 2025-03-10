package pokemoncli

type Pokemon struct {
	Id   int
	Name string
	Img  string
}

type PokemonRepo interface {
	GetPokemons() ([]Pokemon, error)
}

func NewPokemon(id int, name, img string) Pokemon {
	return Pokemon{
		Id:   id,
		Name: name,
		Img:  img,
	}
}
