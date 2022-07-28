package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /
func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}
