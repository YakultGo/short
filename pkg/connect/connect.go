package connect

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

// client 是一个全局的 http.Client，用于发送 http 请求
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: time.Second * 2,
}

// Get 判断给定的 url 是否可以正常访问
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect client.Get failed", logx.LogField{Key: "err", Value: err.Error()})
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
