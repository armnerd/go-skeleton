package mail

import (
	mail "github.com/armnerd/go-skeleton/internal/data/mail"
)

// Add 添加留言
func Add() []mail.Record {
	var model = mail.Record{}
	var data = model.Add()
	return data
}
