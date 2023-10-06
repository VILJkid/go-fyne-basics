package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Application interface {
	Run()
}

type SpeedTestUseCase interface {
	PerformSpeedTest(*widget.Label)
}

type AppBuilder interface {
	BuildUI() fyne.Window
	BuildSpeedTestUseCase() SpeedTestUseCase
}
