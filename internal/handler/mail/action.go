package mail

import (
	"github.com/armnerd/go-skeleton/internal/logic/mail"

	"github.com/gin-gonic/gin"
)

// Add 添加留言
func Add(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	email := c.DefaultPostForm("email", "")
	message := c.DefaultPostForm("message", "")
	mail.Add(name, email, message)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "ok",
		"data":    "",
	})
}
