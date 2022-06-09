# go-skeleton

> Golang 脚手架，Go 简单，Go 直接

## Run

```bash
mkdir runtime
go build main.go

// 获取运行参数
./main server - h

// 开启接口服务
./main server

// 运行脚本
./main cmd demo hello
./main cmd demo world
```

## Lib

|   Role   |   Package   |   Link   |
| ---- | ---- | ---- |
|   命令   |   cobra       |   https://github.com/spf13/cobra     |
|   路由   |   Gin       |   https://github.com/gin-gonic/gin     |
|   配置   |   godotenv  |   https://github.com/joho/godotenv     |
|   ORM    |   Gorm      |   https://github.com/jinzhu/gorm       |
|   redis  |   redigo    |   https://github.com/gomodule/redigo   |
|   Curl   |   goz       |   https://github.com/idoubi/goz        |
|   Json   |   gjson     |   https://github.com/tidwall/gjson     |
|   日志   |   logrus    |   https://github.com/sirupsen/logrus   |
|   鉴权   |   jwt-go    |   https://github.com/dgrijalva/jwt-go  |
|   Kafka   |   kafka-go    |   https://github.com/segmentio/kafka-go  |
|   RabbitMq   |   amqp    |   https://github.com/streadway/amqp  |
|   Etcd   |   etcd    |   https://github.com/coreos/etcd/clientv3 |
