package app

import (
	"fmt"

	"github.com/bigusbeckus/quorum-challenge-backend/internal/pkg/config"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	config := config.AppConfig.Service
	port := fmt.Sprintf(":%d", config.Port)

	if config.GinReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)
	SetupRoutes(router)

	router.Run(port)
}

func SetupRoutes(r *gin.Engine) {
	r.GET("/", RootHandler)
	r.GET("/search", SearchHandler)
}
