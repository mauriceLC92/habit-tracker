package habit

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	Id uuid.UUID
	// update this to be title
	Title       string
	Completed   bool
	DateStarted time.Time
	UserId      uuid.UUID
}

type HabitCreator struct {
	NewId       uuid.UUID
	DateStarted time.Time
	HabitName   string
	UserId      uuid.UUID
}

func CreateHabit(hc HabitCreator) Habit {
	return Habit{
		Id:          hc.NewId,
		Title:       hc.HabitName,
		Completed:   true,
		DateStarted: hc.DateStarted,
		UserId:      hc.UserId,
	}
}
