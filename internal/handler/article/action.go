package article

import (
	"math"
	"strconv"

	"github.com/armnerd/go-skeleton/internal/logic/article"
	"github.com/armnerd/go-skeleton/pkg/response"

	"github.com/gin-gonic/gin"
)

// @Summary 文章列表
// @Produce  json
// @Param page query int false "页码"
// @Param category query int false "分类"
// @Param timeline query int false "时间轴"
// @Param search query string false "搜索"
// @Success 200 {object} listResult "成功"
// @Router /api/article/list [post]
func List(c *gin.Context) {
	// 获取参数
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	category, _ := strconv.Atoi(c.DefaultPostForm("category", ""))
	timeline, _ := strconv.Atoi(c.DefaultPostForm("timeline", ""))
	search := c.DefaultPostForm("search", "")

	// 获取列表
	start := (page - 1) * 6
	var list, err = article.List(c, start, category, timeline, search)
	if err != nil {
		response.Fail(c, response.InternalError)
		return
	}

	// 获取总页数
	var total float64 = 0
	var num = article.Total(c, category, timeline, search)
	if num != 0 {
		total = math.Ceil(float64(num) / 6)
	}

	// 返回数据
	data := listResult{
		Category: category,
		Timeline: timeline,
		Search:   search,
		List:     list,
		Page:     page,
		Total:    total,
		Count:    num,
	}
	response.Succuss(c, data)
}

// @Summary 文章详情
// @Produce  json
// @Param id query int true "文章 id"
// @Success 200 {object} data.Article "成功"
// @Router /api/article/info [post]
func Info(c *gin.Context) {
	// 参数验证
	id := c.DefaultPostForm("id", "")
	if id == "" {
		response.Fail(c, response.ParamsLost)
		return
	}

	// 获取详情
	data, err := article.Info(c, id)
	if err != nil {
		response.Fail(c, response.RecordNotExsit)
		return
	}

	// 返回数据
	response.Succuss(c, data)
}

// Add 新增文章
func Add(c *gin.Context) {
	// 返回数据
	response.Succuss(c, "")
}

// Edit 编辑文章
func Edit(c *gin.Context) {
	// 返回数据
	response.Fail(c, response.ParamsLost)
}
