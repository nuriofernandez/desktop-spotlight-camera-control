package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
)

var deviceUniqueId = os.Getenv("DEVICE_UNIQUE_ID")
var bearerToken = os.Getenv("HOMEBRIDGE_TOKEN")

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

	// Authenticate request
	req.Header.Add("Authorization", "Bearer "+bearerToken)
	req.Header.Add("Content-Type", "application/json")

	// Perform request
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	// No error, function end
	return nil
}
