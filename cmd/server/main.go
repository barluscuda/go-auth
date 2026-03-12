package main

import (
	"fmt"
	"log"
	"os"

	"github.com/barluscuda/go-auth/Internal/adapter/router"
	"github.com/barluscuda/go-auth/config"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("server is starting...")
	log.Println("server is running in PID:", os.Getpid())

	cfg, err := config.Load()
	if err != nil {
		log.Panic("failed to load config: " + err.Error())
	}

	{
		httpServer := SetupHttpServer()
		log.Println("http server port:", cfg.ServerPort)
		httpServer.Run(fmt.Sprintf(":%d", cfg.ServerPort))
	}
}

func SetupHttpServer() *gin.Engine {
	cfg, err := config.Load()
	if err != nil {
		log.Panic("failed to load config: " + err.Error())
	}

	switch cfg.ServerMode {
	case "dev":
		log.Println("server starting in development mode")
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		log.Panicf("unknown server mode: %q", cfg.ServerMode)
	}

	serve := gin.Default()
	router.SetupRoute(serve)

	return serve
}
