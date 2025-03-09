package repository

import (
	"bufio"
	musclescli "intro/internal"
	"os"
	"strconv"
	"strings"
)

type repository struct {
}

func NewRepository() musclescli.MuscleRepo {
	return &repository{}
}

func (r *repository) GetMuscles() ([]musclescli.Muscle, error) {
	f, _ := os.Open("/Users/loddrik/Documents/documentos_linux/Curso Golang/introduction/01/data/muscles.csv")
	reader := bufio.NewReader(f)

	var muscles []musclescli.Muscle

	for line := readLine(reader); line != nil; line = readLine(reader) {
		lineValues := strings.Split(string(line), ",")
		idNumber, _ := strconv.Atoi(lineValues[0])

		muscle := musclescli.NewMuscle(
			idNumber,
			lineValues[1],
			lineValues[2],
			lineValues[3],
			musclescli.NewDifficultyType(lineValues[4]),
		)

		muscles = append(muscles, muscle)

	}
	return muscles, nil

}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()

	return
}
