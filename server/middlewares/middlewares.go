package middlewares

import (
	"log/slog"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func CorsMiddleware(maxAge int) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           time.Duration(maxAge) * time.Second,
	})
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		sep := " | "

		slog.InfoContext(c,
			"| "+c.Request.Method+sep+
				strings.Split(c.Request.RequestURI, "?")[0]+sep+
				cast.ToString(c.Writer.Status())+sep+
				cast.ToString(duration)+"s",
		)
	}
}
