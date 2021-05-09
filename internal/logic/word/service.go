package word

import (
	word "github.com/armnerd/go-skeleton/internal/data/word"
)

// List 单词列表
func List() []word.Record {
	var model = word.Record{}
	var data = model.GetAll()
	return data
}

// Info 单词详情
func Info() []word.Record {
	var model = word.Record{}
	var data = model.GetOne()
	return data
}
