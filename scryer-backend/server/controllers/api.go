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
		fmt.Println("Failed to unmarshal OneStepGPS json:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	fmt.Println("Finishing OneStepGpsData action", jsonResponse)
	c.JSON(http.StatusOK, gin.H{
		"response": jsonResponse,
	})
}
