package musclescli

// Muscle is a struct that represents a muscle

type Muscle struct {
	Id         int
	Name       string
	Type       string
	Exercise   string
	Difficulty *DifficultyType
}
type DifficultyType int

const (
	Unknown DifficultyType = iota
	Begginer
	Intermediate
	Advanced
)

func (d DifficultyType) String() string {
	return toString[d]
}

func NewDifficultyType(d string) *DifficultyType {
	var difficultyType DifficultyType

	if _, ok := toID[d]; ok {
		difficultyType = toID[d]
	}
	return &difficultyType
}

var toString = map[DifficultyType]string{
	Unknown:      "Unknown",
	Begginer:     "Begginer",
	Intermediate: "Intermediate",
	Advanced:     "Advanced",
}

var toID = map[string]DifficultyType{
	"Unknown":      Unknown,
	"Begginer":     Begginer,
	"Intermediate": Intermediate,
	"Advanced":     Advanced,
}

type MuscleRepo interface {
	GetMuscles() ([]Muscle, error)
}

func NewMuscle(id int, name string, muscleType string, exercise string, difficulty *DifficultyType) (m Muscle) {
	m = Muscle{
		Id:         id,
		Name:       name,
		Type:       muscleType,
		Exercise:   exercise,
		Difficulty: difficulty,
	}
	return
}
