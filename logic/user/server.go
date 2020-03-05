package user

import (
	. "goto/data"
)

// Hello 打个招呼
func Hello() string {
	var message = "Hi~"
	return message
}

// List 用户列表
func List() []User {
	var User = User{}
	var data = User.GetAll()
	return data
}

// One 一个用户
func One() string {
	var message = "One~"
	return message
}

// Add 新增用户
func Add() string {
	var message = "Add~"
	return message
}

// Delete 删除用户
func Delete() string {
	var message = "Delete~"
	return message
}
