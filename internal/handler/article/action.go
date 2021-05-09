package article

import (
	article "github.com/armnerd/go-skeleton/internal/logic/article"
	response "github.com/armnerd/go-skeleton/pkg/response"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math"
	"strconv"
	"strings"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"

	"github.com/gin-gonic/gin"
)

// List 文章列表
func List(c *gin.Context) {
	// 获取参数
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	category, _ := strconv.Atoi(c.DefaultPostForm("category", ""))
	timeline, _ := strconv.Atoi(c.DefaultPostForm("timeline", ""))
	search := c.DefaultPostForm("search", "")

	// 获取列表
	start := (page - 1) * 6
	var list = article.List(start, category, timeline, search)

	// 获取总页数
	var total float64 = 0
	var num = article.Total(category, timeline, search)
	if num != 0 {
		total = math.Ceil(num / 6)
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

// Info 文章详情
func Info(c *gin.Context) {
	// 参数验证
	id := c.DefaultPostForm("id", "")
	if id == "" {
		response.Fail(c, response.ParamsLost)
		return
	}

	// 获取详情
	var data = article.Info(id)
	if data.ID == 0 {
		response.Fail(c, response.RecordNotExsit)
		return
	}

	// 返回数据
	response.Succuss(c, data)
}

// Add 添加文章
func Add(c *gin.Context) {
	// 返回数据
	response.Succuss(c, "")
}

// Edit 编辑文章
func Edit(c *gin.Context) {
	// 返回数据
	response.Fail(c, response.ParamsLost)
}

// Sync 同步
func Sync(c *gin.Context) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var all = article.FetchAll()
	for _, article := range all {
		jsonStr, _ := json.Marshal(article)
		iterm := string(jsonStr)

		req := esapi.IndexRequest{
			Index:      "article",
			DocumentID: strconv.Itoa(article.ID),
			Body:       strings.NewReader(iterm),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("[%s] Error indexing document ID=%d", res.Status(), article.ID)
		} else {
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				log.Printf("Error parsing the response body: %s", err)
			} else {
				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			}
		}
	}

	response.Succuss(c, all)
}

// Search 全文检索
func Search(c *gin.Context) {
	search := c.DefaultPostForm("search", "")
	if search == "" {
		response.Fail(c, response.ParamsLost)
		return
	}
	var r map[string]interface{}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  search,
				"fields": [2]string{"Title", "Raw"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("article"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	var items []map[string]interface{}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		one := hit.(map[string]interface{})["_source"].(map[string]interface{})
		score := hit.(map[string]interface{})["_score"]

		countSplit := strings.Split(one["Raw"].(string), " ")
		result := ""
		for _, line := range countSplit {
			inline := strings.Contains(line, search)
			if inline {
				line = strings.Replace(line, "```", "", -1)
				line = strings.Replace(line, "\n", "", -1)
				result += strings.Replace(line, search, " <mark>"+search+"</mark> ", -1)
			}
		}

		instance := map[string]interface{}{
			"ID":    one["ID"],
			"Title": one["Title"],
			"Hit":   result,
			"Score": score,
		}

		items = append(items, instance)
	}

	// 返回数据
	response.Succuss(c, items)
}
