package router

import (
	"airbnb-auth-be/internal/pkg/env"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DefaultCORSSetting() gin.HandlerFunc {

	return cors.New(cors.Config{
		AllowOrigins: env.CONFIG.Origins,
		AllowMethods: []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "Cache-Control", "X-Requested-With"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func NewRouter() *gin.Engine {
	if env.CONFIG.Stage == string(env.StageLocal) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	return router
}
