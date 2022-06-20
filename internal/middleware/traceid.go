package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get("X-Request-Id")
		if traceId == "" {
			u4 := uuid.New()
			traceId = u4.String()
		}
		c.Set("traceId", traceId)
		c.Writer.Header().Set("X-Request-Id", traceId)
		c.Next()
	}
}
