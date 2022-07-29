package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"scryer-backend/db/models"
)

var IdentityKey = "ID"

type userLoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	// TODO: Put secret key in ENV
	return jwt.New(&jwt.GinJWTMiddleware{
		Key: []byte(jwtSecretKey),
		Timeout: time.Hour,
		MaxRefresh: time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.ID,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			user := models.User{ID: uint(claims[IdentityKey].(float64))}
			user.FindByID()
			return &user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var input userLoginForm

			if err := c.ShouldBindJSON(&input); err != nil {
				fmt.Println(input)
				panic(err.Error())
				fmt.Println(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()

				return nil, jwt.ErrFailedAuthentication
			}

			user := models.User{Username: input.Username, Password: input.Password}
			if err := user.HashPassword(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()

				return nil, jwt.ErrFailedAuthentication
			}

			if err := user.FindByUsername(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()

				return nil, jwt.ErrFailedAuthentication
			}
			fmt.Println("Authenticated, userId")
			fmt.Println(user.ID)
			return &user, nil
		},
	})
}
