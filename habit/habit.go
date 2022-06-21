package habit

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	Id          string
	Name        string
	Completed   bool
	DateStarted time.Time
	UserId      int
}

type HabitCreator struct {
	NewId       uuid.UUID
	DateStarted time.Time
	HabitName   string
	UserId      int
}

func CreateHabit(hc HabitCreator) Habit {
	return Habit{
		Id:          hc.NewId.String(),
		Name:        hc.HabitName,
		Completed:   true,
		DateStarted: hc.DateStarted,
		UserId:      hc.UserId,
	}
}
