package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/jkhf/go_project/internal/app"
	"github.com/jkhf/go_project/internal/routes"
)

// panic is a self-destructive function that stops the program and prints the error message
func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf(("Runing app on port %d\n"), port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
