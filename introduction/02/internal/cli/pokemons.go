package cli

import (
	"fmt"
	"os"
	pokemoncli "pokemon-cli/internal"

	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, arg []string)

const (
	idFlag   = "id"
	csvFlag  = "csv"
	fileFlag = "file"
)

func InitPokemonCmd(repository pokemoncli.PokemonRepo) *cobra.Command {
	pokemonCmd := &cobra.Command{
		Use:   "pokemon",
		Short: "Prints a list of pokemons",
		Run:   runPokemonsFn(repository),
	}

	// Add all flags
	pokemonCmd.Flags().IntP(idFlag, "i", 0, "Pokemon ID")
	pokemonCmd.Flags().BoolP(csvFlag, "c", false, "Export pokemons to CSV")
	pokemonCmd.Flags().StringP(fileFlag, "f", "", "CSV file name")

	return pokemonCmd
}

func runPokemonsFn(repository pokemoncli.PokemonRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt(idFlag)
		exportCSV, _ := cmd.Flags().GetBool(csvFlag)
		fileName, _ := cmd.Flags().GetString(fileFlag)

		pokemons, err := repository.GetPokemons()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting pokemons: %v\n", err)
			return
		}

		if exportCSV {
			if fileName == "" {
				fmt.Println("Please provide a file name for the CSV export.")
				return
			}
			repository.DumpPokemonsToCSV(pokemons, fileName)
			return
		}

		// Handle ID filter if specified
		if id != 0 {
			displayPokemonByID(pokemons, id)
			return
		}

		// Default: display all pokemons
		displayAllPokemons(pokemons)
	}
}

func displayPokemonByID(pokemons []pokemoncli.Pokemon, id int) {
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			fmt.Println("===================================")
			fmt.Printf("Pokemon: %s\n", pokemon.Name)
			fmt.Printf("ID: %d\n", pokemon.Id)
			fmt.Printf("Image URL: %s\n", pokemon.Img)
			fmt.Println("===================================")
			return
		}
	}
	fmt.Println("Pokemon not found with that ID")
}

func displayAllPokemons(pokemons []pokemoncli.Pokemon) {
	fmt.Println("All Pokemons:")
	fmt.Println("===================================")
	for _, pokemon := range pokemons {
		fmt.Printf("ID: %d - Name: %s\n", pokemon.Id, pokemon.Name)
	}
	fmt.Println("===================================")
}
