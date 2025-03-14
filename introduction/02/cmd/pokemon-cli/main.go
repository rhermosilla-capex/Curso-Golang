package main

import (
	"pokemon-cli/internal/cli"
	"pokemon-cli/internal/storage/pokeapi"

	"github.com/spf13/cobra"
)

func main() {

	// csvRepo := csv.NewRepository()
	apiRepo := pokeapi.NewPokeapiRepository()

	rootCmd := &cobra.Command{Use: "pokemon-cli"}
	rootCmd.AddCommand(cli.InitPokemonCmd(apiRepo))
	rootCmd.Execute()
}
