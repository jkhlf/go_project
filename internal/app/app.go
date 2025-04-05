package app

//housing = data

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jkhf/go_project/internal/api"
	"github.com/jkhf/go_project/internal/store"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

// AWS Secrets for DB in real Production
// Pointer to the Application and error type, nil is error type
func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//stores will go here

	//handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is avaible\n")
}
