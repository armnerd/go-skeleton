package redis

import (
	"testing"

	"github.com/armnerd/go-skeleton/pkg/redis"
	"github.com/smartystreets/goconvey/convey"
)

func TestSet(t *testing.T) {
	convey.Convey("测试从 redis 写入数据", t, func() {
		redis.GetCache()
		defer func() {
			redis.Get().Close()
		}()
		key := "redis_test"
		value := "hello"
		_, err := redis.Get().Do("SET", key, value)
		convey.So(err, convey.ShouldBeNil)
	})
}

func TestGett(t *testing.T) {
	convey.Convey("测试从 redis 读取数据", t, func() {
		redis.GetCache()
		defer func() {
			redis.Get().Close()
		}()
		key := "redis_test"
		res, err := redis.Get().Do("GET", key)
		convey.So(err, convey.ShouldBeNil)
		convey.So(res, convey.ShouldResemble, []byte("hello"))
	})
}
