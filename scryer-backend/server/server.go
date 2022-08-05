package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"scryer-backend/db"
	"scryer-backend/server/controllers"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("ID")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not logged in"})
		c.Abort()
		return
	}
}

func Run(dbContext *db.DbContext) {
	// wrap controllers that need database with DbContext
	authController := controllers.AuthController{DbContext: dbContext}
	userController := controllers.UserController{DbContext: dbContext}

	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")

	r.Use(sessions.Sessions("session", cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))))

	r.GET("/", controllers.Index)

	// Use a simple cache for OneStepGPS API so I'm not spamming their API too much
	store := persistence.NewInMemoryStore(time.Second)
	r.GET("/ping", cache.CachePage(store, time.Minute, controllers.OneStepGpsData))

	r.POST("/register", userController.CreateUser)
	r.POST("/login", authController.Login)

	// Both these routes require login
	private := r.Group("/")
	private.Use(AuthRequired)
	private.GET("/logout", authController.Logout)
	private.PUT("/user/preferences", userController.UpdateDevicePreferences)

	r.Run()
}
