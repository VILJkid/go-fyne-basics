// window.go
package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/VILJkid/go-fyne-basics/app"
)

type AppBuilder struct {
	app         fyne.App
	updateChan  chan *fyne.Container // Channel to update the window content
	CurrentView *fyne.Container      // Current view container
}

func NewAppBuilder(app fyne.App, updateChan chan *fyne.Container) *AppBuilder {
	return &AppBuilder{
		app:        app,
		updateChan: updateChan,
	}
}

func (ab *AppBuilder) BuildUI() fyne.Window {
	window := ab.app.NewWindow("Fyne Basics | Go")

	// Set the size of the window and make it fixed
	window.Resize(fyne.NewSize(720, 480))
	window.SetFixedSize(true)

	label := widget.NewLabel("Click the Go button to start the speed tests.")
	label.Alignment = fyne.TextAlignCenter

	btn := widget.NewButton("Go", func() {
		// Trigger the speed test
		fmt.Println("Go")
		label = widget.NewLabel("Searching closest servers...")
		ab.updateView(container.NewVBox(
			label,
		))
		go func() {
			// Perform the speed test
			resultLabel := widget.NewLabel("")
			ab.BuildSpeedTestUseCase().PerformSpeedTest(resultLabel)
			ab.updateView(container.NewVBox(
				resultLabel,
			))
			fmt.Println("Done")
		}()
	})

	resultBox := widget.NewLabel("")
	content := container.NewVBox(
		label,
		btn,
		container.NewVBox(resultBox),
	)
	ab.CurrentView = content

	window.SetContent(content)

	return window
}

func (ab *AppBuilder) BuildSpeedTestUseCase() app.SpeedTestUseCase {
	// Create and return the SpeedTestUseCase
	return app.NewSpeedTestInteractor()
}

func (ab *AppBuilder) updateView(newView *fyne.Container) {
	// Send the new view to the update channel
	ab.updateChan <- newView
}
