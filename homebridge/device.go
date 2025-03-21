package homebridge

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var deviceUniqueId = os.Getenv("DEVICE_UNIQUE_ID")

func TurnSpotlightTo(status bool) error {
	// Build data
	url := "http://homebridge.local/api/accessories/" + deviceUniqueId
	payload := strings.NewReader("{\"characteristicType\": \"On\",\"value\": \"" + strconv.FormatBool(status) + "\"}")

	// Prepare request
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, payload)

	if err != nil {
		return err
	}

	// Required content type
	req.Header.Add("Content-Type", "application/json")

	// Authenticate request
	token, err := getAccessToken()

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	// Perform request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// In case of an auth issue, clear the session.
	if res.StatusCode == 401 {
		clearSession() // Clear session in case it's expired
		return errors.New(res.Status)
	}

	// No error, function end
	return nil
}
