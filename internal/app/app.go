package app

//housing = data

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

// Pointer to the Application and error type, nil is error type
func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
