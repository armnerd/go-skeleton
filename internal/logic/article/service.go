package article

import (
	"github.com/armnerd/go-skeleton/internal/data/article"
)

// List 文章列表
func List(start int, category int, timeline int, search string) []article.Record {
	var model = article.Record{}
	var data = model.GetAll(start, category, timeline, search)
	return data
}

// FetchAll 所有文章
func FetchAll() []article.Record {
	var model = article.Record{}
	var data = model.FetchAll()
	return data
}

// Total 文章总数
func Total(category int, timeline int, search string) float64 {
	var model = article.Record{}
	var data = model.GetTotal(category, timeline, search)
	return data
}

// Info 文章列表
func Info(id string) article.Record {
	var model = article.Record{}
	var data = model.GetOne(id)
	return data
}
