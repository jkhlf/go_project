package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/jkhf/go_project/internal/app"
	"github.com/jkhf/go_project/internal/routes"
)

//go get -> bring the package
// go install -> same npm install

// panic is a self-destructive function that stops the program and prints the error message
func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("we are running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
