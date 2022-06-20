package mysql

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB mysql连接池
var DB *gorm.DB

// GetDB 连接数据库
func GetDB() {
	DB = connectDbMySQL(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_CHARSET"),
	)
	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	openConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxOpenConns(maxConnections)
	sqlDB.SetMaxIdleConns(openConnections)
}

// 初始化Mysql db
func connectDbMySQL(host, port, database, user, pass, charset string) *gorm.DB {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: NewTraceLogger(logger.Info, time.Second),
	})
	if err != nil {
		log.Fatalf("models.InitDbMySQL err: %v", err)
	}
	return db
}

// 获取实例
func Instance(c *gin.Context) *gorm.DB {
	return DB.WithContext(Gin2Context(c))
}

// 追加 traceId
func Gin2Context(c *gin.Context) context.Context {
	return context.WithValue(context.Background(), "traceId", c.GetString("traceId"))
}
