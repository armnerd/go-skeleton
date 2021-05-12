package mail

import (
	"time"

	"github.com/armnerd/go-skeleton/pkg/mysql"
)

// Add 添加留言
func (*Record) Add(param map[string]string) bool {
	data := Record{
		Name:    param["name"],
		Email:   param["email"],
		Message: param["message"],
		Ctime:   int(time.Now().Unix()),
	}
	mysql.DB.Create(&data)
	return true
}
