package app

import (
	"fyne.io/fyne/v2"
)

type AppInitializer struct {
	app     fyne.App
	builder AppBuilder
}

func NewAppInitializer(app fyne.App, builder AppBuilder) *AppInitializer {
	return &AppInitializer{
		app:     app,
		builder: builder,
	}
}

func (ai *AppInitializer) Initialize() {
	// Create a new window and set the title
	window := ai.builder.BuildUI()

	// Run the application
	window.ShowAndRun()
}

func (ai *AppInitializer) GetApp() fyne.App {
	return ai.app
}
