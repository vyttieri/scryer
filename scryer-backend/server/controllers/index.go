package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /
// Unused in dev environment since front-end is being served from Vite
func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}
