package middleware

import (
	"time"
	"zychimne/instant/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(c config.CorsConfig) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     c.AllowMethods,
		AllowHeaders:     c.AllowHeaders,
		AllowCredentials: c.AllowCreds,
		AllowOrigins:     c.AllowOrigins,
		MaxAge:           time.Duration(c.MaxAge) * time.Hour,
	})
}
