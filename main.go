package main

import (
	_ "github.com/joho/godotenv/autoload" // Important to keep this as the first import!

	"github.com/nuriofernandez/desktop-spotlight-camera-control/homebridge"
	"time"
)

func main() {
	for {
		cameraStatus := checkCameraStatus()
		_ = homebridge.TurnSpotlightTo(cameraStatus)

		time.Sleep(1 * time.Second)
	}
}
