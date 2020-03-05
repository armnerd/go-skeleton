package main

import (
	mysql "goto/database/mysql"
	route "goto/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	application := gin.Default()
	application = route.Init(application)
	mysql.ConnectDB()
	application.Run()
}
