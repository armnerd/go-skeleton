package response

import (
	"github.com/gin-gonic/gin"
)

// Succuss 成功
func Succuss(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
		"traceid": c.GetString("traceId"),
	})
}

// Fail 失败
func Fail(c *gin.Context, code int) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": GetMessageByCode(code),
		"data":    "",
		"traceid": c.GetString("traceId"),
	})
}
