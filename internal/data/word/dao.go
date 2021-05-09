package word

import "github.com/armnerd/go-skeleton/pkg/mysql"

// GetAll 分页获取单词列表
func (*Record) GetAll() []Record {
	var data = make([]Record, 0)
	mysql.DB.Find(&data)
	return data
}

// GetOne 获取单词详情
func (*Record) GetOne() []Record {
	var data = make([]Record, 0)
	mysql.DB.Find(&data)
	return data
}
