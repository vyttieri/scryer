package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"scryer-backend/db/models"
)

type userLoginForm struct {
	Username string `json:"username" binding:"required,gte=1,lte=32"`
	Password string `json:"password" binding:"required,gte=1,lte=64"`
}

func Login(c *gin.Context) {
	fmt.Println("Logging in user")

	var input userLoginForm
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Failed to bind login input", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()

		return
	}

	user := models.User{Username: input.Username, Password: input.Password}
	if err := user.HashPassword(); err != nil {
		fmt.Println("Failed to hash password", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()

		return
	}

	if err := user.FindByUsername(); err != nil {
		fmt.Println("Failed to find user", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()

		return
	}

	session := sessions.Default(c)
	session.Set("ID", user.ID)
	if err := session.Save(); err != nil {
		fmt.Println("Failed to save session", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()

		return
	}

	// Is there a way to do this in one fell swoop?
	var devicePreferences []models.DevicePreference
	if err := user.FindDevicePreferences(&devicePreferences); err != nil {
		fmt.Println("Failed to find devicePreferences", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()

		return
	}
	user.DevicePreferences = devicePreferences

	fmt.Println("Successfully logged in, userId:", user.ID)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Logout(c *gin.Context) {
	fmt.Println("Logging out user")

	session := sessions.Default(c)
	session.Delete("ID")
	if err := session.Save(); err != nil {
		fmt.Println("Failed to logout", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()

		return
	}

	fmt.Println("successfully logged out user")

	c.JSON(http.StatusOK, gin.H{})
}
