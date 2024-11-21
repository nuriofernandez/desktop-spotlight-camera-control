package homebridge

import "time"

var homebridgeSessionExpire time.Time
var homebridgeToken = ""

func getAccessToken() (string, error) {
	// Check if session expires in this minute
	currentTime := time.Now().Add(time.Minute)
	if currentTime.After(homebridgeSessionExpire) {
		homebridgeToken = ""
	}

	// Use current session token
	if homebridgeToken != "" {
		return homebridgeToken, nil
	}

	// Generate a new session
	session, err := homebridgeAuth()
	if err != nil {
		return "", err
	}

	// Store new session details
	homebridgeSessionExpire = time.Now().Add(time.Duration(session.ExpiresIn) * time.Second)
	homebridgeToken = session.AccessToken

	// Return the stored session
	return homebridgeToken, nil
}
