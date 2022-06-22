package habit_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/mauriceLC92/habit-tracker/habit"
)

const (
	TestUuid     string = "10d75b29-01e3-4b47-9ae7-2036c146119c"
	TestUserUuid string = "67f4380b-4a51-4a69-bbee-66bd9fd10856"
)

func TestHabit(t *testing.T) {
	t.Parallel()
	_ = habit.Habit{
		Id:          uuid.MustParse(TestUuid),
		Title:       "running",
		Completed:   false,
		DateStarted: time.Now(),
	}
}

func fakeHabitCreator(t *testing.T) habit.HabitCreator {
	testDate := time.Date(2022, 03, 3, 4, 33, 3, 3, time.UTC)
	return habit.HabitCreator{
		NewId:       uuid.MustParse(TestUuid),
		HabitName:   "running",
		UserId:      uuid.MustParse(TestUserUuid),
		DateStarted: testDate,
	}
}

func TestCreateHabit(t *testing.T) {
	t.Parallel()

	hc := fakeHabitCreator(t)
	want := habit.Habit{
		Id:          uuid.MustParse(TestUuid),
		Title:       hc.HabitName,
		Completed:   true,
		DateStarted: hc.DateStarted,
		UserId:      hc.UserId,
	}

	got := habit.CreateHabit(hc)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetHabitById(t *testing.T) {
	t.Parallel()

}
