package router

import (
	"github.com/barluscuda/go-auth/Internal/adapter/handler"
	"github.com/barluscuda/go-auth/Internal/adapter/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoute(serve *gin.Engine) {
	serve.NoRoute(middleware.NoRoute)
	api := serve.Group("/api")
	{
		api.GET("ping", handler.ApiPing)
	}
}