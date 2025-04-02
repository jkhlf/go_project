package routes

import (
	"github.com/go-chi/chi"
	"github.com/jkhf/go_project/internal/app"
)

//curl localhost:8080/workouts/2
//curl localhost:8080/workouts

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHandler.HandleGetWorkout)
	return r
}
