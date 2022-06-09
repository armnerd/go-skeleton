package curl

import (
	"testing"

	"github.com/armnerd/go-skeleton/pkg/curl"
	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("测试 get 请求", t, func() {
		welcome := "This is go-skeleton, build with Gin and Gorm"
		var url = "http://127.0.0.1:9551"
		data := map[string]interface{}{}
		header := map[string]interface{}{}
		content, err := curl.Get(url, data, header)
		convey.So(err, convey.ShouldBeNil)
		res := content.Get("Welcome").Value()
		convey.So(res, convey.ShouldEqual, welcome)
	})
}

func TestPost(t *testing.T) {
	convey.Convey("测试 post 请求", t, func() {
		var url = "http://127.0.0.1:9551/api/article/info"
		data := map[string]interface{}{
			"id": "95",
		}
		header := map[string]interface{}{}
		_, err := curl.PostForm(url, data, header)
		convey.So(err, convey.ShouldBeNil)
	})
}
