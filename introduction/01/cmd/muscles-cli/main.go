package main

import (
	muscles "intro/internal/cli"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "muscles-cli"}
	rootCmd.AddCommand(muscles.InitMusclesCmd())
	rootCmd.Execute()
}
