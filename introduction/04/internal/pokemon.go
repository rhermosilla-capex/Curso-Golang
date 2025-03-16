package pokemon

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type AbilityDetail struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}

type Abilities = []AbilityDetail

type PokemonDetail struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Height    int       `json:"height"`
	Weight    int       `json:"weight"`
	Abilities Abilities `json:"abilities"`
}

type PokemonRepo interface {
	GetPokemons() ([]Pokemon, error)
	GetPokemonDetail(id int) (PokemonDetail, error)
}

func NewPokemon(id int, name string) Pokemon {
	return Pokemon{
		Id:   id,
		Name: name,
	}
}

func NewAbility(name, url string) Ability {
	return Ability{
		Name: name,
		Url:  url,
	}
}

func NewAbilityDetail(isHidden bool, slot int, ability Ability) AbilityDetail {
	return AbilityDetail{
		IsHidden: isHidden,
		Slot:     slot,
		Ability:  ability,
	}
}
