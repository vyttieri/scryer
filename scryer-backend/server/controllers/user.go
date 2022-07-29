package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"scryer-backend/db/models"
)

// I really dislike "Preference" vs "Preferences", they should both be "Preferences"!
// But I think it's a necessary evil to distinguish between the singular and the plural.
type DevicePreferenceInput struct {
	DeviceID string `json:"device_id" binding:"required"`
	SortPosition uint `json:"sort_position" binding:"required"`
	Visible bool `json:"visible" binding:"required"`
	Icon string `json:"icon" binding:"required"`
}

type createUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`

	DevicePreferences []DevicePreferenceInput `json:"device_preferences" binding:"required"`
}

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


	c.JSON(http.StatusOK, gin.H{"userId": user.ID, "username": user.Username })
	fmt.Println("Authenticated, userId")
	fmt.Println(user.ID)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("ID")
	if err := session.Save(); err != nil {
		fmt.Println("Failed to save session", err)
		panic(err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{})
}

// POST /users/
func CreateUser(c *gin.Context) {
	var input createUserInput
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

	if err := user.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()

		return
	}

	c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "username": user.Username})
}

// PUT /users/:id/preferences
func UpdateUserPreferences(c *gin.Context) {

}
