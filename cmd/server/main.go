package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/cesarcruzc/go-production-ready-api/internal/transport/http"
)

// App - the struct wich contains things like pointers
// to database connections
type App struct{}

// Run - sets up application
func (app *App) Run() error {
	fmt.Println("Setting up API")
	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8000", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API")
		fmt.Println(err)
	}
}
