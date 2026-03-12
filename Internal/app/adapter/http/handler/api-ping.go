package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ApiPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "pong",
		"server_time": time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
	})
}