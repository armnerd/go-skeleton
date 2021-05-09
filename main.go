package main

import (
	"github.com/armnerd/go-skeleton/pkg/mysql"
	"github.com/armnerd/go-skeleton/pkg/redis"
	"github.com/armnerd/go-skeleton/internal/route"
	"github.com/armnerd/go-skeleton/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 配置
	config.SetAppRoot(os.Args[0])
	godotenv.Load(config.AppRoot + "/.env")
	if os.Getenv("GIN_MODE") == "release" {
		// 生产模式
		gin.SetMode(gin.ReleaseMode)
	}
	// 路由
	app := route.Init()
	// 连接池
	mysql.ConnectDB()
	redis.CreatePool()
	app.Run(":9551")
}
