package mail

import (
	"github.com/armnerd/go-skeleton/pkg/mysql"
)

// Add 添加留言
func (*Record) Add() []Record {
	var data = make([]Record, 0)
	mysql.DB.Find(&data)
	return data
}
