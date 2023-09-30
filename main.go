package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	// Create a new fyne application
	application := app.New()

	// Create a new window and set the title
	window := application.NewWindow("Fyne Basics | Go")

	// Run the application
	window.ShowAndRun()
}
