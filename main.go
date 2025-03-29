package main

import (
	"fmt"
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

	http.HandleFunc("/health", HealthCheck)
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

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is avaible\n")
}

//curl localhost:8080/health
