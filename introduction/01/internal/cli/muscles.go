package muscles

import (
	"fmt"
	musclescli "intro/internal"
	"strconv"

	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

const idFLag = "id"

func InitMusclesCmd(repository musclescli.MuscleRepo) *cobra.Command {
	musclesCmd := &cobra.Command{
		Use:   "muscles",
		Short: "List muscles and their locations",
		Run:   runMusclesFn(repository),
	}
	musclesCmd.Flags().StringP(idFLag, "i", "", "Name of the muscle to get location")
	return musclesCmd
}

func runMusclesFn(repository musclescli.MuscleRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		muscles, _ := repository.GetMuscles()

		id, _ := cmd.Flags().GetString(idFLag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			for _, muscle := range muscles {
				if muscle.Id == i {
					fmt.Println(muscle)
					return
				}
			}
		} else {
			fmt.Print(muscles)

		}

	}
}
