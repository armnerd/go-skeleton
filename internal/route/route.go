package route

import (
	"github.com/armnerd/go-skeleton/internal/handler/article"
	"github.com/armnerd/go-skeleton/internal/handler/mail"
	"github.com/armnerd/go-skeleton/internal/handler/test"
	"github.com/armnerd/go-skeleton/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() *gin.Engine {
	app := gin.New()
	// 中间件
	app.Use(gin.Logger(), middleware.Cors(), middleware.RecoverAtLast())
	// 接口不存在
	app.NoRoute(middleware.NotFound())
	// 路由分组
	api := app.Group("/api/")

	// Welcome
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Welcome": "This is go-skeleton, build with Gin and Gorm",
		})
	})

	// Blog
	api.POST("/article/list", article.List)                            // 文章列表
	api.POST("/article/info", article.Info)                            // 文章详情
	api.POST("/article/add", middleware.AuthRequired(), article.Add)   // 新增文章
	api.POST("/article/edit", middleware.AuthRequired(), article.Edit) // 编辑文章
	api.POST("/article/es/sync", article.Sync)                         // 同步文章
	api.POST("/article/es/search", article.Search)                     // 搜索文章
	api.POST("/feedback", mail.Add)                                    // 留言

	// Redis
	api.POST("/cache/set", test.SetCache)
	api.POST("/cache/get", test.GetCache)

	// Curl
	api.GET("/curl/get", test.CurlGet)
	api.GET("/curl/post", test.CurlPost)

	// Jwt
	api.GET("/jwt/login", test.Login)
	api.GET("/jwt/auth", middleware.AuthRequired(), test.Auth)

	return app
}
