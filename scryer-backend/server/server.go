package server

import (
	// "fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"

	"scryer-backend/server/controllers"
)

func Run() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")

	store := persistence.NewInMemoryStore(time.Second)

	router.GET("/", controllers.Index)

	// TODO: Change cache expiration time
	// TODO: Error handling
	router.GET("/ping", cache.CachePage(store, time.Hour, controllers.OneStepGpsData))

	router.POST("/users/register", controllers.CreateUser)
	router.POST("/users/login", controllers.LoginUser)

	router.Run()
}
