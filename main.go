package main

import (
	_ "github.com/joho/godotenv/autoload" // Important to keep this as the first import!

	"time"
)

func main() {
	for {
		cameraStatus := checkCameraStatus()
		_ = TurnSpotlightTo(cameraStatus)

		time.Sleep(1 * time.Second)
	}
}
