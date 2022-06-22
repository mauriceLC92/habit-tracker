package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mauriceLC92/habit-tracker/habit"
	db "github.com/mauriceLC92/habit-tracker/internal/db/habits"
)

//  todo change this please
var USER_ID string = "95a4d6bd-b811-4885-b5a3-3989feb4306f"

type Greeting struct {
	Message string `json:"message"`
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	greeting := Greeting{
		Message: "Hello from Habitat 🎯",
	}

	json, err := json.Marshal(greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

type CreateHabitRequest struct {
	Name string
}

type CreateHabitResponse struct {
	Id string
}

func (app *application) createHabit(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var h CreateHabitRequest

	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		app.errorLog.Print(err)
	}

	hc := habit.HabitCreator{
		NewId:       uuid.New(),
		DateStarted: time.Now(),
		HabitName:   h.Name,
		UserId:      uuid.MustParse(USER_ID),
	}

	// how to make this better??
	// The habit package should deal with the logic to check if the habit has already been set with 24 hours
	// so in that case we need the habit package to have access to the persistence layer.
	// How would this look when iniatilizing the habit package then?
	// Probably some constructor which you pass in your desired persistence layer (database/file/something)
	// that's then added onto the application struct
	habit := habit.CreateHabit(hc)
	dbHabit, err := app.Queries.CreateHabit(ctx, db.CreateHabitParams{
		ID:     habit.Id,
		Title:  habit.Title,
		UserID: habit.UserId,
	})
	if err != nil {
		app.errorLog.Printf("createHabit: err - %v", err)
	}

	res := CreateHabitResponse{
		Id: dbHabit.ID.String(),
	}

	json, err := json.Marshal(res)
	if err != nil {
		app.errorLog.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
