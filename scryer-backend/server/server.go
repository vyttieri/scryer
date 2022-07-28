package server

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"

	"scryer-backend/db/models"
	"scryer-backend/server/controllers"
)

var identityKey = "ID"

type UserLoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Run() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key: []byte("SECRET KEY"),
		Timeout: time.Hour,
		MaxRefresh: time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{ ID: claims[identityKey].(uint)}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var input UserLoginForm

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

			if err := user.Find(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()

				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},
	})

	if err != nil {
		panic(err)
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		panic(errInit.Error())
	}

	store := persistence.NewInMemoryStore(time.Second)

	r.GET("/", controllers.Index)

	// TODO: Change cache expiration time
	// TODO: Error handling
	r.GET("/ping", cache.CachePage(store, time.Hour, controllers.OneStepGpsData))

	r.POST("/register", controllers.CreateUser)
	r.POST("/login", authMiddleware.LoginHandler)

	r.Run()
}
