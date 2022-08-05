package api

import (
	"fmt"
	"io"
	"net/http"

	"scryer-backend/env"
)

const apiUrl = "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key="
func GetDeviceData() ([]byte, error) {
	var url = fmt.Sprintf("%s%s", apiUrl, env.ApiKey)

	fmt.Println("Starting OneStepGPS GET, URL: %s", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully retrieved OneStepGPS data")
	return body, nil
}
