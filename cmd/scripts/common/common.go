package common

import (
	"os"

	"github.com/armnerd/go-skeleton/config"
	"github.com/armnerd/go-skeleton/pkg/mysql"
	"github.com/armnerd/go-skeleton/pkg/redis"
	"github.com/joho/godotenv"
)

func Depend() {
	// 根目录
	config.SetAppRoot(os.Args[0])
	// 配置
	configFile := config.AppRoot + "/.env"
	godotenv.Load(configFile)
	// 连接池
	mysql.GetDB()
	redis.GetCache()
}

func Release() {
	// 关闭 MySQL
	mysql.DB.Close()
	// 关闭 Redis
	redis.Get().Close()
}
