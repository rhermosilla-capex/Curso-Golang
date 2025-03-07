package muscles

import (
	"fmt"

	"github.com/spf13/cobra"
)

var muscles = map[string]string{
	"biceps":          "front of upper arm",
	"triceps":         "back of upper arm",
	"deltoids":        "shoulder",
	"pectorals":       "chest",
	"abdominals":      "stomach",
	"quadriceps":      "front of thigh",
	"hamstrings":      "back of thigh",
	"gastrocnemius":   "calf",
	"gluteus_maximus": "buttocks",
}

var exercises = map[string]string{
	"biceps":          "bicep curl",
	"triceps":         "tricep dip",
	"deltoids":        "shoulder press",
	"pectorals":       "bench press",
	"abdominals":      "crunch",
	"quadriceps":      "squat",
	"hamstrings":      "deadlift",
	"gastrocnemius":   "calf raise",
	"gluteus_maximus": "hip thrust",
}

type CobraFn func(cmd *cobra.Command, args []string)

const nameFlag = "name"
const exercisesFlag = "exercises"

func InitMusclesCmd() *cobra.Command {
	musclesCmd := &cobra.Command{
		Use:   "muscles",
		Short: "List muscles and their locations",
		Run:   runMusclesFn(),
	}
	musclesCmd.Flags().StringP(nameFlag, "n", "", "Name of the muscle to get location")
	musclesCmd.Flags().StringP(exercisesFlag, "e", "", "Name of the muscle to get exercises")
	return musclesCmd
}

func runMusclesFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(nameFlag)
		exercise, _ := cmd.Flags().GetString(exercisesFlag)

		switch {
		case name != "" && exercise != "":
			printMuscleAndExercise(name)
		case name != "" && exercise == "":
			printMuscleLocation(name)
		case name == "" && exercise != "":
			printExercise(exercise)
		default:
			printAllMusclesAndExercises()
		}
	}
}

func printMuscleAndExercise(name string) {
	fmt.Println("Muscle:", name)
	fmt.Println("Location:", muscles[name])
	fmt.Println("Exercise:", exercises[name])
}

func printMuscleLocation(name string) {
	fmt.Println("Muscle:", name)
	fmt.Println("Location:", muscles[name])
}

func printExercise(exercise string) {
	fmt.Println("Muscle:", exercise)
	fmt.Println("Exercise:", exercises[exercise])
}

func printAllMusclesAndExercises() {
	fmt.Println("Muscles and their locations:")
	for muscle, location := range muscles {
		fmt.Printf("%s: %s\n", muscle, location)
	}
	fmt.Println("\nExercises for each muscle:")
	for muscle, exercise := range exercises {
		fmt.Printf("%s: %s\n", muscle, exercise)
	}
}
