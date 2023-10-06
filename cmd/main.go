// main.go
package main

import (
	"github.com/VILJkid/go-fyne-basics/app"
	"github.com/VILJkid/go-fyne-basics/ui"

	"fyne.io/fyne/v2"
	fApp "fyne.io/fyne/v2/app"
)

func main() {
	myApp := fApp.New()
	updateChan := make(chan *fyne.Container)
	builder := ui.NewAppBuilder(myApp, updateChan)
	appInitializer := app.NewAppInitializer(myApp, builder)

	go func() {
		for {
			// Listen for updates on the content channel and update the window content
			newView := <-updateChan
			builder.CurrentView.Objects = newView.Objects
			builder.CurrentView.Refresh()
		}
	}()

	appInitializer.Initialize()
}
