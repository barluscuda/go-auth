package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "ROUTE NOT FOUND",
		"method": c.Request.Method,
		"path": c.Request.URL.Path,
	})
}