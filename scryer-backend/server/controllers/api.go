package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"scryer-backend/api"
)

// TODO: Error handling
// GET /ping
func OneStepGpsData(context *gin.Context) {
	var response []byte = api.GetDeviceData()

	fmt.Println("Unmarshalling json")
	var jsonResponse interface{} // TODO: Should this be non generic?
	if err := json.Unmarshal(response, &jsonResponse); err != nil {
		panic(err)
	}

	fmt.Println(jsonResponse)
	// TODO: Use Standard JSON Response format
	context.JSON(http.StatusOK, gin.H{
		"response": jsonResponse,
	})
}
