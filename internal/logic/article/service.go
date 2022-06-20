package article

import (
	"github.com/armnerd/go-skeleton/internal/data"
	"github.com/gin-gonic/gin"
)

// List 文章列表
func List(c *gin.Context, start int, category int, timeline int, search string) ([]data.Article, error) {
	var model = data.Article{}
	return model.GetAll(c, start, category, timeline, search)
}

// Total 文章总数
func Total(c *gin.Context, category int, timeline int, search string) int64 {
	var model = data.Article{}
	return model.GetTotal(c, category, timeline, search)
}

// Info 文章详情
func Info(c *gin.Context, id string) (data.Article, error) {
	var model = data.Article{}
	return model.GetOne(c, id)
}
