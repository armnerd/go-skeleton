# GoTo

Golang脚手架，Go简单，Go直接

## DataFlow

```
main => route => handler => logic => data
```

## Undo

```
JWT
中间件[跨域|鉴权]
redigo
```

## Lib

|   Role   |   Package   |   Link   |
| ---- | ---- | ---- |
|   框架   |   Gin   |   https://github.com/gin-gonic/gin   |
|   ORM   |   Gorm   |   https://github.com/jinzhu/gorm   |
|   配置   |   godotenv   |   https://github.com/joho/godotenv   |


## Nginx

```
server
{
    listen 80;
    server_name goto;

    location /{
         root /Users/zane/go/src/goto/dist;
    }
    
    location /api/ {
         proxy_pass http://127.0.0.1:8080;
    }
}
```


## Sql

```
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `mobile` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1
```


## Reference

```
https://github.com/qq1060656096/gapp.git
https://github.com/710leo/Toruk
https://github.com/bydmm/singo
```

