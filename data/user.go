package data

import (
	"goto/database/mysql"
)

// User 用户表
type User struct {
	ID     int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"type:varchar(45)"`
	Mobile string `gorm:"type:varchar(45)"`
}

// GetAll 获取所有用户
func (*User) GetAll() []User {
	var users = make([]User, 0)
	mysql.DB.Find(&users)
	return users
}
