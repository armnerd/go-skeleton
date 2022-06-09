package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

// redis连接池
var pool *redis.Pool

// GetCache 获取缓存池
func GetCache() {
	host := os.Getenv("CACHE_HOST")
	port := os.Getenv("CACHE_PORT")
	pass := os.Getenv("CACHE_PASS")
	server := fmt.Sprintf("%s:%s", host, port)
	pool = &redis.Pool{
		MaxActive:   100,
		MaxIdle:     100,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			// 鉴权
			if pass != "" {
				if _, err := c.Do("AUTH", pass); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
}

// Get 获取实例
func Get() redis.Conn {
	return pool.Get()
}
