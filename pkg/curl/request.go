package curl

import (
	syslog "github.com/armnerd/go-skeleton/pkg/log"

	"github.com/idoubi/goz"
	"github.com/tidwall/gjson"
)

// Get 请求
func Get(url string, data map[string]interface{}, header map[string]interface{}) (gjson.Result, error) {
	var res gjson.Result
	cli := goz.NewClient()
	resp, err := cli.Get(url, goz.Options{
		Headers: header,
		Query: data,
	})
	if err != nil {
		syslog.Error("http_get", err.Error())
		return res, err
	}
	body, err := resp.GetBody()
	if err != nil {
		syslog.Error("http_get", err.Error())
		return res, err
	}
	res = gjson.Parse(body.GetContents())
	return res, nil
}

// Post 请求
func Post(url string, data map[string]interface{}, header map[string]interface{}) (gjson.Result, error) {
	var res gjson.Result
	cli := goz.NewClient()
	resp, err := cli.Post(url, goz.Options{
		Headers: header,
		FormParams: data,
	})
	if err != nil {
		syslog.Error("http_post", err.Error())
		return res, err
	}
	body, err := resp.GetBody()
	if err != nil {
		syslog.Error("http_post", err.Error())
		return res, err
	}
	res = gjson.Parse(body.GetContents())
	return res, nil
}
