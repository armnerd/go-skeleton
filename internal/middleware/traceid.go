package middleware

import (
	"github.com/gin-gonic/gin"
	//"github.com/satori/go.uuid"
	"github.com/google/uuid"
)

func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestId := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestId == "" {
			u4 := uuid.New()
			//u4 := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use in the application
		c.Set("traceId", requestId)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
