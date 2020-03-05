package user

import (
	user "goto/logic/user"

	"github.com/gin-gonic/gin"
)

// Hello 打个招呼
func Hello(c *gin.Context) {
	var message = user.Hello()
	c.JSON(200, gin.H{
		"code":    0,
		"message": message,
		"data":    "",
	})
}

// List 用户列表
func List(c *gin.Context) {
	var data = user.List()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
	})
}

// One 一个用户
func One(c *gin.Context) {
	var message = user.One()
	c.JSON(200, gin.H{
		"code":    0,
		"message": message,
		"data":    "",
	})
}

// Add 新增用户
func Add(c *gin.Context) {
	var message = user.Add()
	c.JSON(200, gin.H{
		"code":    0,
		"message": message,
		"data":    "",
	})
}

// Delete 删除用户
func Delete(c *gin.Context) {
	var message = user.Delete()
	c.JSON(200, gin.H{
		"code":    0,
		"message": message,
		"data":    "",
	})
}
