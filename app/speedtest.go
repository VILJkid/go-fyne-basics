package app

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
	"github.com/showwin/speedtest-go/speedtest"
)

type SpeedTestInteractor struct {
	label *widget.Label
}

func NewSpeedTestInteractor() *SpeedTestInteractor {
	return &SpeedTestInteractor{}
}

func (st *SpeedTestInteractor) PerformSpeedTest(label *widget.Label) {
	st.label = label
	// Create a new Speedtest client
	speedClient := speedtest.New()

	st.label.SetText("Searching closest servers...")
	st.label.Refresh()

	// Get the closest server
	serverList, _ := speedClient.FetchServers()
	targets, _ := serverList.FindServer(nil)
	server := targets[0]

	tests := []string{"Ping", "Download", "Upload"}
	for _, test := range tests {
		switch test {
		case "Ping":
			st.label.SetText(test + " Test in progress...")
			st.label.Refresh()
			server.PingTest(nil)
		case "Download":
			st.label.SetText(test + " Test in progress...")
			st.label.Refresh()
			server.DownloadTest()
		case "Upload":
			st.label.SetText(test + " Test in progress...")
			st.label.Refresh()
			server.UploadTest()
		}
	}
	var result string
	result += fmt.Sprintf("Ping: %d ms\n", server.Latency.Milliseconds())
	result += fmt.Sprintf("Download: %.2f Mbps\n", server.DLSpeed)
	result += fmt.Sprintf("Upload: %.2f Mbps", server.ULSpeed)

	// Update the label with the test results
	st.label.SetText(result)
	st.label.Refresh()
}
