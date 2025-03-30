package api

import (
	"net/http"

	"fmt"

	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Get workout by ID: ", workoutID)
}

func (wh *WorkoutHandler) HandleGetWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "created a workout\n")
}
