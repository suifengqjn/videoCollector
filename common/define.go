package common

import "errors"

const (
	DownloadTimeFormat = "2006-01-02"
	PlatZY = "zuiyou"

)

var (
	ProxyError = errors.New("ProxyError")

)


func BrowserHeader() map[string]string  {
	header := make(map[string]string)
	header["Connection"]="keep-alive"
	header["User-Agent"]="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML]=like Gecko) Chrome/77.0.3865.90 Safari/537.36"
	header["Accept"]="text/html,application/xhtml+xml,application/xml;q=0.9,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"
	header["Accept-Language"]="zh-CN,zh;q=0.9,en;q=0.8"
	return header
}