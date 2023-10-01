package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	// Create a new fyne application
	application := app.New()

	// Create a new window and set the title
	window := application.NewWindow("Fyne Basics | Go")

	// Set the size of the window and
	window.Resize(fyne.NewSize(720, 480))
	window.SetFixedSize(true)

	var content *fyne.Container
	var label *widget.Label
	var btn, okButton *widget.Button
	// var initializeContent func()

	initializeContent := func() {
		fmt.Println("Inside initializeContent")
		// Create a new label
		label = widget.NewLabel("Click the Go button to start the speed tests.")
		label.Alignment = fyne.TextAlignCenter

		btn = widget.NewButton("Go", func() {
			fmt.Println("Go Button clicked")

			// Show a loading spinner
			loadingSpinner := widget.NewProgressBarInfinite()

			content = container.NewVBox(
				label,
				loadingSpinner,
			)
			window.SetContent(content)

			// Perform three individual speed tests one by one
			go triggerSpeedTest(window, content, label, loadingSpinner, okButton)
		})

		resultBox := widget.NewLabel("Hehe")

		content = container.NewVBox(
			label,
			btn,
			container.NewVBox(resultBox),
		)
	}

	initializeContent()

	// Set the content of the window
	window.SetContent(content)

	// Run the application
	window.ShowAndRun()
}

func triggerSpeedTest(w fyne.Window, c *fyne.Container, label *widget.Label, loadingSpinner *widget.ProgressBarInfinite, okButton *widget.Button) {
	performSpeedTest(w, label)
	c.Remove(loadingSpinner)
	c.Refresh()

	// Add an "OK" button after displaying results
	okButton = widget.NewButton("OK", clickOK)
	c.Add(okButton)
}

func clickOK() {
	fmt.Println("OK button cliked")
	// Reset the content to the initial state
	// initializeContent()
}

func performSpeedTest(w fyne.Window, label *widget.Label) {
	// Create a new Speedtest client
	speedClient := speedtest.New()

	// Get the closest server
	label.SetText("Searching closest servers...")

	serverList, _ := speedClient.FetchServers()
	targets, _ := serverList.FindServer(nil)
	server := targets[0]

	tests := []string{"Ping", "Download", "Upload"}
	for _, test := range tests {
		switch test {
		case "Ping":
			label.SetText(test + " Test in progress...")
			server.PingTest(nil)
		case "Download":
			label.SetText(test + " Test in progress...")
			server.DownloadTest()
		case "Upload":
			label.SetText(test + " Test in progress...")
			server.UploadTest()
		}
	}
	var result string
	result += fmt.Sprintf("Ping: %d ms\n", server.Latency.Milliseconds())
	result += fmt.Sprintf("Download: %.2f Mbps\n", server.DLSpeed)
	result += fmt.Sprintf("Upload: %.2f Mbps", server.ULSpeed)

	label.SetText(result)
}
