package main

import (
	"github.com/jkhf/go_project/internal/app"
)

// panic is a self-destructive function that stops the program and prints the error message
func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
}
