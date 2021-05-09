package article

import (
	"github.com/armnerd/go-skeleton/pkg/mysql"
)

// GetAll 分页获取文章列表
func (*Record) GetAll(start int, category int, timeline int, search string) []Record {
	var data = make([]Record, 0)

	// 参数绑定
	var params = &Record{}
	if category != 0 {
		params.Type = category
	}
	if timeline != 0 {
		params.Timeline = timeline
	}
	tx := mysql.DB.Where(params)

	// 搜索参数
	if search != "" {
		tx = tx.Where("title LIKE ?", "%"+search+"%")
	}
	tx.Offset(start).Limit(6).Order("id desc").Find(&data)

	return data
}

// FetchAll 获取所有文章
func (*Record) FetchAll() []Record {
	var data = make([]Record, 0)
	mysql.DB.Order("id desc").Find(&data)

	return data
}

// GetTotal 获取文章总数
func (*Record) GetTotal(category int, timeline int, search string) float64 {
	var count float64

	// 参数绑定
	var params = &Record{}
	if category != 0 {
		params.Type = category
	}
	if timeline != 0 {
		params.Timeline = timeline
	}
	tx := mysql.DB.Model(&Record{}).Where(params)

	// 搜索参数
	if search != "" {
		tx = tx.Where("title LIKE ?", "%"+search+"%")
	}
	tx.Count(&count)

	return count
}

// GetOne 获取文章详情
func (*Record) GetOne(id string) Record {
	var data = Record{}
	mysql.DB.First(&data, id)
	return data
}
