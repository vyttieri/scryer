package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"scryer-backend/db/models"
)

type UserCreateForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}

// POST /users/register
func CreateUser(c *gin.Context) {
	var input UserCreateForm
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

// POST /users/login
func LoginUser(c *gin.Context) {

}

// GET /users/:id/preferences
func GetUserPreferences(c *gin.Context) {

}

// POST /users/:id/preferences
func CreateUserPreferences(c *gin.Context) {

}

// PUT /users/:id/preferences
func UpdateUserPreferences(c *gin.Context) {

}
