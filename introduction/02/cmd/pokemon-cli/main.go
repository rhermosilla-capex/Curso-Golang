package main

import "github.com/spf13/cobra"

func main() {

	// csvRepo := csv.NewRepository()
	// apiRepo := api.NewRepository()

	rootCmd := &cobra.Command{Use: "pokemon-cli"}
	// rootCmd.AddCommand(cli.InitMusclesCmd(csvRepo))
	rootCmd.Execute()
}
