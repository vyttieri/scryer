package server

import (
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"

	"scryer-backend/db/models"
	"scryer-backend/server/auth"
	"scryer-backend/server/controllers"
)


func Run() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")

	authMiddleware, err := auth.CreateAuthMiddleware()

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

	auth := r.Group("/auth")

	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	r.Run()
}

func helloHandler(c *gin.Context) {
	fmt.Println("hit helloHandler")
  claims := jwt.ExtractClaims(c)
  user, _ := c.Get(auth.IdentityKey)
  c.JSON(200, gin.H{
    "userID":   claims[auth.IdentityKey],
    "username": user.(*models.User).Username,
    "text":     "Hello World.",
  })
}
