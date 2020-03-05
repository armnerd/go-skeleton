package route

import (
	user "goto/handler/user"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(app *gin.Engine) *gin.Engine {
	application := gin.Default()

	api := application.Group("/api/")

	// 增删改查
	api.GET("/hi", user.Hello)
	api.GET("/user/list", user.List)
	api.GET("/user/one", user.One)
	api.POST("/user/add", user.Add)
	api.POST("/user/delete", user.Delete)

	return application
}
