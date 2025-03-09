package main

import (
	cli "intro/internal/cli"
	csv "intro/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	csvRepo := csv.NewRepository()


	rootCmd := &cobra.Command{Use: "muscles-cli"}
	rootCmd.AddCommand(cli.InitMusclesCmd(csvRepo))
	rootCmd.Execute()
}
