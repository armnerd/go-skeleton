package data

import (
	"github.com/armnerd/go-skeleton/pkg/mysql"
	"github.com/gin-gonic/gin"
)

// GetAll 分页获取文章列表
func (*Article) GetAll(c *gin.Context, start int, category int, timeline int, search string) ([]Article, error) {
	var data = make([]Article, 0)

	// 参数绑定
	var params = &Article{}
	if category != 0 {
		params.Type = category
	}
	if timeline != 0 {
		params.Timeline = timeline
	}
	tx := mysql.Instance(c).Where(params)

	// 搜索参数
	if search != "" {
		tx = tx.Where("title LIKE ?", "%"+search+"%")
	}
	err := tx.Offset(start).Limit(6).Order("id desc").Find(&data).Error
	return data, err
}

// GetTotal 获取文章总数
func (*Article) GetTotal(c *gin.Context, category int, timeline int, search string) int64 {
	var count int64

	// 参数绑定
	var params = &Article{}
	if category != 0 {
		params.Type = category
	}
	if timeline != 0 {
		params.Timeline = timeline
	}
	tx := mysql.Instance(c).Model(&Article{}).Where(params)

	// 搜索参数
	if search != "" {
		tx = tx.Where("title LIKE ?", "%"+search+"%")
	}
	tx.Count(&count)

	return count
}

// GetOne 获取文章详情
func (*Article) GetOne(c *gin.Context, id string) (Article, error) {
	var data = Article{}
	err := mysql.Instance(c).First(&data, id).Error
	return data, err
}
