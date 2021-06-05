package main

import "fmt"

// App - the struct wich contains things like pointers
// to database connections
type App struct{}

// Run - sets up application
func (app *App) Run() error {
	fmt.Println("Setting up API")
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
