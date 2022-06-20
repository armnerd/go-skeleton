package route

import (
	_ "github.com/armnerd/go-skeleton/docs"
	"github.com/armnerd/go-skeleton/internal/handler/article"
	"github.com/armnerd/go-skeleton/internal/handler/demo"
	"github.com/armnerd/go-skeleton/internal/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() *gin.Engine {
	app := gin.New()
	// 中间件
	app.Use(gin.Logger(), middleware.Cors(), middleware.RecoverAtLast(), middleware.TraceId())
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

	// swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Blog
	api.POST("/article/list", article.List)                            // 文章列表
	api.POST("/article/info", article.Info)                            // 文章详情
	api.POST("/article/add", middleware.AuthRequired(), article.Add)   // 新增文章
	api.POST("/article/edit", middleware.AuthRequired(), article.Edit) // 编辑文章

	// Redis
	api.POST("/cache/set", demo.SetCache)
	api.POST("/cache/get", demo.GetCache)

	// Curl
	api.GET("/curl/get", demo.CurlGet)
	api.GET("/curl/post", demo.CurlPost)

	// Log
	api.GET("/log", demo.Log)

	// Jwt
	api.GET("/jwt/login", demo.Login)
	api.GET("/jwt/auth", middleware.AuthRequired(), demo.Auth)

	return app
}
