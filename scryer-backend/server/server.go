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

	// Initialize JWT authentication module
	authMiddleware, err := auth.CreateAuthMiddleware()
	if err != nil {
		panic(err)
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		panic(errInit.Error())
	}

	r.GET("/", controllers.Index)

	// Use a simple cache for OneStepGPS API so I'm not spamming their API too much
	// Timeout can be tailored based on use case.
	store := persistence.NewInMemoryStore(time.Second)
	// TODO: Change cache expiration time
	// TODO: Error handling
	r.GET("/ping", cache.CachePage(store, time.Hour, controllers.OneStepGpsData))

	r.POST("/users", controllers.CreateUser)
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
  fmt.Println(user)
  c.JSON(200, gin.H{
    "userID":   claims[auth.IdentityKey],
    "username": user.(*models.User).Username,
    "text":     "Hello World.",
  })
}
