package mail

import (
	mail "github.com/armnerd/go-skeleton/internal/logic/mail"

	"github.com/gin-gonic/gin"
)

// Add 添加留言
func Add(c *gin.Context) {
	var res = mail.Add()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    res,
	})
}
