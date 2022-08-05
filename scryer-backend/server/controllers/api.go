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

	response, err := api.GetDeviceData()
	if err != nil {
		fmt.Println("Error fetching OneStepGPS data: %s", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
		return
	}

	fmt.Println("Unmarshalling json from OneStepGPS")
	var jsonResponse interface{}
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
