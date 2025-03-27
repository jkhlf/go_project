package main

import (
	"net/http"
	"time"

	"github.com/jkhf/go_project/internal/app"
)

// panic is a self-destructive function that stops the program and prints the error message
func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println(("Runing app"))

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
