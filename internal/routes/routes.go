package routes

import (
	"github.com/go-chi/chi"
	"github.com/jkhf/go_project/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	return r
}
