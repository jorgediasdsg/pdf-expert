package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jorgediasdsg/pdf-expert/internal/log"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()
		start := time.Now()

		c.Set("request_id", reqID)

		log.Logger.Info("request_start",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"req_id", reqID,
		)

		c.Next()

		log.Logger.Info("request_end",
			"status", c.Writer.Status(),
			"duration_ms", time.Since(start).Milliseconds(),
			"req_id", reqID,
		)
	}
}
