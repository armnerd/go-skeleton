package word

import (
	word "github.com/armnerd/go-skeleton/internal/logic/word"

	"github.com/gin-gonic/gin"
)

// List 单词列表
func List(c *gin.Context) {
	var data = word.List()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
	})
}

// Info 单词详情
func Info(c *gin.Context) {
	var data = word.Info()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
	})
}
