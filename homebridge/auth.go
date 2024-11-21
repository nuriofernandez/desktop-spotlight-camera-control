package homebridge

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

var homebridgeUser = os.Getenv("HOMEBRIDGE_USER")
var homebridgePass = os.Getenv("HOMEBRIDGE_PASSWORD")

type HomebridgeSession struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func homebridgeAuth() (s HomebridgeSession, err error) {
	url := "http://homebridge.local/api/auth/login"
	method := "POST"

	payload := strings.NewReader(`{"username":"` + homebridgeUser + `","password":"` + homebridgePass + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return s, err
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return s, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return s, err
	}

	// Extract 'access_token'
	json.Unmarshal(body, &s)
	return s, nil
}
