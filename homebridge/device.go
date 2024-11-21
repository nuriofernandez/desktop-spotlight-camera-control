package homebridge

import (
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
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	// No error, function end
	return nil
}
