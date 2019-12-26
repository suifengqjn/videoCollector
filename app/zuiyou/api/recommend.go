package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myProject/videoCollector/app/zuiyou/model"
	"myProject/videoCollector/common"
	"myTool/other"
	"net/http"
	"strings"
)

//推荐
func GetRecommend() []*common.VideoModel {
	param:="{\"direction\":\"down\",\"h_model\":\"iPhone SE\",\"h_ch\":\"appstore\",\"h_ua\":\"Mozilla\\/5.0 (iPhone; CPU iPhone OS 12_4_1 like Mac OS X) AppleWebKit\\/605.1.15 (KHTML, like Gecko) Mobile\\/15E148 zuiyou_speed\\/1.0.8\",\"h_app\":\"zuiyou_speed\",\"ad_wakeup\":1,\"c_types\":[1,7,8,11,20,21,22],\"h_nt\":1,\"h_av\":\"1.0.8\",\"tab\":\"rec\",\"h_did\":\"3865a9635d0ed344d4a3b478cbeeae0422f8281e\",\"filter\":\"all\",\"h_os\":1240000,\"h_ts\":1572222222222,\"h_m\":1298829,\"token\":\"T3KdNKOaZ14_QhiQ7KpL9TpjEQSxczv1fDBWrFLrF7utWQtmGserNvU2q8sRl9XLJC3vO\",\"h_idfa\":\"436C228C-2B0B-42EE-A40B-B8B87CAE1EF7\",\"h_dt\":1,\"sdk_ver\":{\"tt_aid\":\"5020119\",\"tx_aid\":\"1109416283\",\"tx\":\"4.7.6\",\"tt\":\"1.9.4\"}}"
	param = strings.ReplaceAll(param, "1572222222222", GetHts())
	sign := other.ZYSign(param)
	url := fmt.Sprintf("https://api.izuiyou.com/index/recommend?sign=%s", sign)

	payload := strings.NewReader(param)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Host", "api.izuiyou.com")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "aliyungf_tc=AQAAADeK9VmwdgIAgkNWy2NghWWZUz4z")
	req.Header.Add("ZYP", "mid=1298829")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "zuiyou_speed/1.0.8 (iPhone; iOS 12.4.1; Scale/2.00)")
	req.Header.Add("Accept-Language", "zh-Hans-CN;q=1, en-US;q=0.9, zh-Hant-CN;q=0.8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var res model.RecommendRes

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil
	}

	if res.Ret == 1 {
		return ParseVideoDetail(res.Data.List)
	} else {
		return nil
	}

}
