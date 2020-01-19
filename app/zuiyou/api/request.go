package api

import (
	"myProject/videoCollector/common"
	"net/http"
)

func ConfigHeader(req *http.Request)  {
	req.Host = "api.izuiyou.com"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", common.ReadConfig().Zy.Cookie)
	req.Header.Set("Zyp", "mid=1298829")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "zuiyou_speed/1.0.8 (iPhone; iOS 12.4.1; Scale/2.00)")
	req.Header.Set("Accept-Language", "zh-Hans-CN;q=1, en-US;q=0.9, zh-Hant-CN;q=0.8")
}
