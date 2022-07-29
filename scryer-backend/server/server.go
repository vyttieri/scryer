package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"scryer-backend/server/controllers"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()

		panic("User is not logged in")
	}
}

func Run() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")

	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("SECRETMOVEELSEWHERE"))))

	router.GET("/", controllers.Index)

	// Use a simple cache for OneStepGPS API so I'm not spamming their API too much
	// Timeout can be tailored based on use case.
	store := persistence.NewInMemoryStore(time.Second)
	// TODO: Change cache expiration time
	// TODO: Error handling
	router.GET("/ping", cache.CachePage(store, time.Hour, controllers.OneStepGpsData))

	private := router.Group("/")
	private.Use(AuthRequired)


	router.POST("/users", controllers.CreateUser)
	router.POST("/login", controllers.Login)

	router.Run()
}
