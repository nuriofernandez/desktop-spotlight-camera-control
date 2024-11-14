package main

import mediadevicesstate "github.com/antonfisher/go-media-devices-state"

func checkCameraStatus() bool {
	// Check camera status
	isCameraOn, err := mediadevicesstate.IsCameraOn()

	// If something goes wrong, return false
	if err != nil {
		return false
	}

	// Return camera status
	return isCameraOn
}
