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

//关注
func GetAttentionUp() []*common.VideoModel {

	param := "{\"direction\":\"down\",\"h_model\":\"iPhone SE\",\"h_ch\":\"appstore\",\"h_app\":\"zuiyou_speed\",\"up_offset\":0,\"c_types\":[1],\"h_ts\":1572222222222,\"h_av\":\"1.0.8\",\"h_nt\":1,\"h_did\":\"3865a9635d0ed344d4a3b478cbeeae0422f8281e\",\"h_m\":1298829,\"h_os\":1240000,\"token\":\"T8K4NKOaZ14_QhiQ7KpL9TpjEQXYhfTjLNOM1eoHOF_QaYOfpCTOUo1gBcN8mwKrIWcc9\",\"h_dt\":1,\"down_offset\":0}"
	param = strings.ReplaceAll(param, "1572222222222", GetHts())
	sign := other.ZYSign(param)
	url := fmt.Sprintf("https://api.izuiyou.com/attention/follow_list?sign=%s", sign)

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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var res model.AttentionRes

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
