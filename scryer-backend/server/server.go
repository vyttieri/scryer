package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"

	"scryer-backend/api"
)

func Run() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")

	store := persistence.NewInMemoryStore(time.Second)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Posts",
		})
	})

	// TODO: Change cache expiration time
	// TODO: Error handling
	router.GET("/ping", cache.CachePage(store, time.Hour, func(c *gin.Context) {
		var response []byte = api.GetDeviceData()

		fmt.Println("Unmarshalling json")
		var jsonResponse interface{} // TODO: Should this be non generic?
		if err := json.Unmarshal(response, &jsonResponse); err != nil {
			panic(err)
		}

		fmt.Println(jsonResponse)
		// TODO: Use Standard JSON Response format
		c.JSON(http.StatusOK, gin.H{
			"response": jsonResponse,
		})
	}))

	router.Run()
}
