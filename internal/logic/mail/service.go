package mail

import (
	"github.com/armnerd/go-skeleton/internal/data/mail"
)

// Add 添加留言
func Add(name string, email string, message string) bool {
	var model = mail.Record{}
	var param = map[string]string{
		"name":    name,
		"email":   email,
		"message": message,
	}
	return model.Add(param)
}
