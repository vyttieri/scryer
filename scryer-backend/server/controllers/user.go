package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"scryer-backend/db"
	"scryer-backend/db/models"
)

type UserController struct {
	DbContext *db.DbContext
}

type createUserInput struct {
	Username string `json:"username" binding:"required,gte=1,lte=32"`
	Password string `json:"password" binding:"required,gte=1,lte=64"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required,gte=1,lte=64"`

	DevicePreferenceInputs []DevicePreferenceInput `json:"devicePreferences" binding:"required"`
}

// I really dislike "Preference" vs "Preferences", they should both be "Preferences"!
// But I think it's a necessary evil to distinguish between the singular and the plural.
type DevicePreferenceInput struct {
	DeviceID string `json:"deviceId" binding:"required"`
	SortPosition uint `json:"sortPosition" binding:"required"`
	Visible *bool `json:"visible" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}

// POST /register
func (controller *UserController) CreateUser(c *gin.Context) {
	fmt.Println("Started registering new user")

	var input createUserInput
	// Ensure we have all the required fields
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Failed to bind createUserInput", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	user := models.User{Username: input.Username, Password: input.Password}
	if err := user.HashPassword(); err != nil {
		fmt.Println("Failed to hash password", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate DevicePreferences
	devicePreferences := make([]models.DevicePreference, len(input.DevicePreferenceInputs))
	for i, DevicePreferenceInput := range input.DevicePreferenceInputs {
		devicePreferences[i] = models.DevicePreference{
			DeviceID: DevicePreferenceInput.DeviceID,
			Icon: DevicePreferenceInput.Icon,
			SortPosition: DevicePreferenceInput.SortPosition,
			Visible: *DevicePreferenceInput.Visible,
		}
	}
	user.DevicePreferences = devicePreferences
	// Create User & DevicePreferences in DB
	if err := user.Create(controller.DbContext); err != nil {
		fmt.Println("Failed to create User & DevicePreferences", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Create our session
	session := sessions.Default(c)
	session.Set("ID", user.ID)
	if err := session.Save(); err != nil {
		fmt.Println("Failed to save user session", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Good job team
	fmt.Println("registered user")
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

type updateDevicePreferencesInput struct {
	DevicePreferenceInputs []UpdateDevicePreferenceInput `json:"devicePreferences" binding:"required"`
}

// Since we are updating, we should have the ID
type UpdateDevicePreferenceInput struct {
	ID uint `json:"id" binding:"required"`
	DeviceID string `json:"deviceId" binding:"required"`
	SortPosition uint `json:"sortPosition" binding:"required"`
	Visible *bool `json:"visible" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}

// PUT /user/preferences
func (controller *UserController) UpdateDevicePreferences(c *gin.Context) {
	fmt.Println("Starting update device preferences")

	var input updateDevicePreferencesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Failed to bind updateDevicePreferencesInput", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	UserID := session.Get("ID").(uint)
	user := models.User{ID: UserID}
	if err := user.FindByID(controller.DbContext); err != nil {
		fmt.Println("Failed to find user", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	devicePreferences := make([]models.DevicePreference, len(input.DevicePreferenceInputs))
	for i, devicePreferenceInput := range input.DevicePreferenceInputs {
		devicePreferences[i] = models.DevicePreference{
			ID: devicePreferenceInput.ID,
			Icon: devicePreferenceInput.Icon,
			SortPosition: devicePreferenceInput.SortPosition,
			Visible: *devicePreferenceInput.Visible,
		}
	}
	if err := user.UpdateDevicePreferences(controller.DbContext, &devicePreferences); err != nil {
		fmt.Println("Failed to update device preferences in database", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	fmt.Println("Successfully updated device preferences")
	c.JSON(http.StatusOK, gin.H{})
}
