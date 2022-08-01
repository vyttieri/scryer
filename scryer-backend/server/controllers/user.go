package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"scryer-backend/db/models"
)

type userLoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input userLoginForm
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(input)
		panic(err.Error())
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()

		return
	}

	user := models.User{Username: input.Username, Password: input.Password}
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()

		return
	}

	if err := user.FindByUsername(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()

		return
	}
	session := sessions.Default(c)
	session.Set("ID", user.ID)
	if err := session.Save(); err != nil {
		fmt.Println("Failed to login", err)
		panic(err)
	}

	// Is there a way to do this in one fell swoop?
	var devicePreferences []models.DevicePreference
	if err := user.FindDevicePreferences(&devicePreferences); err != nil {
		panic(err)
		return
	}
	user.DevicePreferences = devicePreferences

	c.JSON(http.StatusOK, gin.H{"user": user})
	fmt.Println("Authenticated, userId")
	fmt.Println(user.ID)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("ID")
	if err := session.Save(); err != nil {
		fmt.Println("Failed to save session", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// I really dislike "Preference" vs "Preferences", they should both be "Preferences"!
// But I think it's a necessary evil to distinguish between the singular and the plural.
type DevicePreferenceInput struct {
	DeviceID string `json:"deviceId" binding:"required"`
	SortPosition uint `json:"sortPosition" binding:"required"`
	Visible *bool `json:"visible" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}

type createUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`

	DevicePreferenceInputs []DevicePreferenceInput `json:"devicePreferences" binding:"required"`
}

// POST /users/
func CreateUser(c *gin.Context) {
	var input createUserInput
	// Ensure we have all the required fields
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(input)
		panic(err.Error())
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()

		return
	}
	fmt.Println("after binding, input in CreateUser:")
	fmt.Println(input)
	// Hash password
	user := models.User{Username: input.Username, Password: input.Password}
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()

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
	fmt.Println("generated devicePreferences: ", devicePreferences)
	// Create User & DevicePreferences in DB
	if err := user.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()

		return
	}

	// Create our session
	session := sessions.Default(c)
	session.Set("ID", user.ID)
	if err := session.Save(); err != nil {
		fmt.Println("Failed to login", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

		return
	}

	// Good job team
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

// PUT /users/:id/preferences
func UpdateDevicePreferences(c *gin.Context) {
	var input updateDevicePreferencesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(input)
		panic(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	session := sessions.Default(c)
	UserID := session.Get("ID").(uint)
	user := models.User{ID: UserID}
	if err := user.FindByID(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		c.Abort()

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

		fmt.Println("devicePreference visible", *devicePreferenceInput.Visible)
	}
	if err := user.UpdateDevicePreferences(&devicePreferences); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		c.Abort()

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
