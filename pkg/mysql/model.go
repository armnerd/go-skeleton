package mysql

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // indirect
)

// DB mysql连接池
var DB *gorm.DB

// ConnectDB 连接数据库
func ConnectDB() {
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
	DB.DB().SetMaxOpenConns(maxConnections)
	DB.DB().SetMaxIdleConns(openConnections)
	dbLog := os.Getenv("DB_LOG")
	if dbLog == "true" {
		DB.LogMode(true)
	}
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

	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Fatalf("models.InitDbMySQL err: %v", err)
	}
	db.SingularTable(true)
	return db
}
