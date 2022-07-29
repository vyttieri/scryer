package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var apiKey = os.Getenv("ONESTEPGPS_API_KEY")
var url = fmt.Sprintf("https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=%s", apiKey)

func GetDeviceData() []byte {
	fmt.Println("starting onestepgps GET")
	fmt.Printf("getting URL %s", url)
	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	fmt.Println("Response status:", response.Status)

	defer response.Body.Close()

	body, error := io.ReadAll(response.Body)
	fmt.Println(string(body))
//	fmt.Printf("body: %s", body)
	// fmt.Println(json.Marshal(error))
	return body
}

// TODO: Only send data to front-end that it needs.
// Error frontend: device.latest_accurate_device_point.lat_lng "latest_accurate_device_point" is null
// should probably display updated_at on the frontend
// One time