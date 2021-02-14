package middleware

import (
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	localHost = "http://localhost:3000"
)

var (
	corsOnce       sync.Once
	corsMiddleware gin.HandlerFunc
)

// CorsMiddleware return the middeware instance
func CorsMiddleware() gin.HandlerFunc {
	corsOnce.Do(func() {
		corsMiddleware = cors.New(cors.Config{
			AllowOrigins: []string{localHost},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders: []string{
				"Content-Length",
				"Origin",
				"cookie",
				"authorization",
				"origin",
				"content-type",
				"Content-Type",
				"accept",
				"X-CSRF-Token",
				"x-requested-with",
				"Cache-Control",
			},
			ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})
	})
	return corsMiddleware
}
