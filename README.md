# go-skeleton

> Golang 脚手架，Go 简单，Go 直接

## Run

```bash
// 配置
cp .env.example .env

// 日志
mkdir runtime

// 构建
go build main.go

// 构建时指定 CPU 核心数量
go build -ldflags "-X main.SetCpuCount=1" main.go

// 获取运行参数
./main server -h

// 开启接口服务
./main server

// 运行脚本
./main cmd demo hello
./main cmd demo world
```

## Doc

* 依赖

```bash
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0 
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

* 初始化

```bash
swag init
```

* 访问

http://127.0.0.1:9551/swagger/index.html

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

## Undo

* 链路追踪 jaeger / zipkin
* 单元测试 test
