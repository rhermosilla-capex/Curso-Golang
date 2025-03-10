package cli

import (
	"fmt"
	pokemoncli "pokemon-cli/internal"

	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, arg []string)

const idFlag = "id"

func InitPokemonCmd(repository pokemoncli.PokemonRepo) *cobra.Command {
	pokemonCmd := &cobra.Command{
		Use:   "pokemon",
		Short: "Prints a list of pokemons",
		Run:   runPokemonsFn(repository),
	}

	pokemonCmd.Flags().IntP(idFlag, "i", 0, "Pokemon ID")
	return pokemonCmd
}

func runPokemonsFn(repository pokemoncli.PokemonRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt(idFlag)
		pokemons, _ := repository.GetPokemons()

		if id != 0 {
			for _, pokemon := range pokemons {
				if pokemon.Id == id {
					fmt.Println(pokemon)
					return
				}
			}
		} else {
			fmt.Println(pokemons)
		}

	}
}
