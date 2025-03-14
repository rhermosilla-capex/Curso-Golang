package csv

import (
	"bufio"
	"os"
	pokemoncli "pokemon-cli/internal"
	"strconv"
	"strings"
)

type repository struct{}

func NewRepository() pokemoncli.PokemonRepo {
	return &repository{}
}

func (r *repository) GetPokemons() ([]pokemoncli.Pokemon, error) {
	f, _ := os.Open("/Users/loddrik/Documents/documentos_linux/Curso Golang/introduction/02/data/pokemons.csv")
	reader := bufio.NewReader(f)

	var pokemons []pokemoncli.Pokemon

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		idNumber, _ := strconv.Atoi(values[0])

		beer := pokemoncli.NewPokemon(
			idNumber,
			values[1],
			values[2],
		)

		pokemons = append(pokemons, beer)

	}

	return pokemons, nil
}

func (r *repository) DumpPokemonsToCSV(pokemons []pokemoncli.Pokemon, fileName string) {
	// Not implemented
}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
