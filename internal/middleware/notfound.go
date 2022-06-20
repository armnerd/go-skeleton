package middleware

import "github.com/gin-gonic/gin"

// NotFound 接口不存在
func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code":    4000,
			"message": "api not found",
			"data":    "",
		})
	}
}
