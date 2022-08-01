package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"scryer-backend/api"
)

// GET /ping
func OneStepGpsData(c *gin.Context) {
	fmt.Println("Starting OneStepGpsData action")
	var response []byte = api.GetDeviceData()

	fmt.Println("Unmarshalling json from OneStepGPS")
	var jsonResponse interface{} // TODO: Should this be non generic?
	if err := json.Unmarshal(response, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()

		return
	}

	fmt.Println("Finishing OneStepGpsData action", jsonResponse)
	// TODO: Use Standard JSON Response format
	c.JSON(http.StatusOK, gin.H{
		"response": jsonResponse,
	})
}
