package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myProject/videoCollector/app/zuiyou/model"
	"myProject/videoCollector/common"
	"net/http"
	"time"
)

func GetFavorList() []*common.VideoModel {

	var res []*common.VideoModel
	group, err := GetFavorGroupList()
	if err != nil {
		return nil
	}

	for _, l := range group.Data.List {

		for _, g := range common.ReadConfig().Zy.Favor.Group {
			if l.Name == g {

				for i:=0;i<10;i++ {
					vs, err := GetFavorDetail(l.ID, i)
					if err != nil || len(vs) == 0 {
						break
					}
					res = append(res, vs...)

				}

			}
		}

	}

	return res
}

// 收藏
/*
{
	"h_model": "iPhone SE",
	"h_ch": "appstore",
	"h_app": "zuiyou_speed",
	"c_types": [1],
	"h_ts": 1569163748588,
	"h_av": "1.0.8",
	"h_nt": 1,
	"h_did": "3865a9635d0ed344d4a3b478cbeeae0422f8281e",
	"t": 0,
	"h_m": 1298829,
	"h_os": 1240000,
	"token": "T3K1NKOaZ14_QhiQ7KpL9TpjEQQrcm8nuAfDb80_l42mmiMS0c9F9oGokpDXXHCZhv4oO",
	"h_dt": 1
}


{
	"ret": 1,
	"data": {
		"list": [{
			"id": 14496105,
			"mid": 1298829,
			"name": "素材",
			"ct": 1564475177,
			"ut": 1564475177,
			"post_count": 147
		}, {
			"id": 11025102,
			"mid": 1298829,
			"name": "动漫",
			"ct": 1542118477,
			"ut": 1542118477,
			"post_count": 2
		}],
		"more": 0,
		"t": 1477323680
	}
}
*/

type FavorListReq struct {
	HModel string `json:"h_model"`
	HCh    string `json:"h_ch"`
	HApp   string `json:"h_app"`
	CTypes []int  `json:"c_types"`
	HTs    int64  `json:"h_ts"`
	HAv    string `json:"h_av"`
	HNt    int    `json:"h_nt"`
	HDid   string `json:"h_did"`
	T      int    `json:"t"`
	HM     int    `json:"h_m"`
	HOs    int    `json:"h_os"`
	Token  string `json:"token"`
	HDt    int    `json:"h_dt"`
}

type FavorListRes struct {
	Ret  int `json:"ret"`
	Data struct {
		List []struct {
			ID        int    `json:"id"`
			Mid       int    `json:"mid"`
			Name      string `json:"name"`
			Ct        int    `json:"ct"`
			Ut        int    `json:"ut"`
			PostCount int    `json:"post_count"`
		} `json:"list"`
		More int `json:"more"`
		T    int `json:"t"`
	} `json:"data"`
}

func GetFavorGroupList() (*FavorListRes, error) {

	data := FavorListReq{
		HModel: "iPhone SE",
		HCh:    "appstore",
		HApp:   "zuiyou_speed",
		CTypes: []int{1},
		HTs:    time.Now().UnixNano() / 1000000,
		HAv:    "1.0.8",
		HNt:    1,
		HDid:   "3865a9635d0ed344d4a3b478cbeeae0422f8281e",
		T:      0,
		HM:     1298829,
		HOs:    1240000,
		Token:  common.ReadConfig().Zy.Token,
		HDt:    1,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.izuiyou.com/favor/list?sign=256486683b422a65a2006b57584c84ae", body)
	if err != nil {
		return nil, err
	}

	ConfigHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var res FavorListRes
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

/*
{
	"id": 14496105,
	"h_model": "iPhone SE",
	"h_ch": "appstore",
	"h_app": "zuiyou_speed",
	"c_types": [1],
	"h_ts": 1569163855335,
	"h_av": "1.0.8",
	"h_nt": 1,
	"h_did": "3865a9635d0ed344d4a3b478cbeeae0422f8281e",
	"offset": 0,
	"h_os": 1240000,
	"h_m": 1298829,
	"token": "T3K1NKOaZ14_QhiQ7KpL9TpjEQQrcm8nuAfDb80_l42mmiMS0c9F9oGokpDXXHCZhv4oO",
	"h_dt": 1
}

{
  "ret": 1,
  "data": {
    "list": [
      {
        "type": 1,
        "data": {
          "topic": {
            "id": 109086,
            "topic": "电影推荐",
            "atts_title": "影迷"
          },
          "videos": {
            "825204537": {
              "dur": 703,
              "thumb": 825204537,
              "playcnt": 23854,
              "url": "http://dlvideo.izuiyou.com/zyvd/03/ff/c1f4-dc72-11e9-af2c-00163e042306",
              "priority": 1,
              "urlsrc": "http://dlvideo.izuiyou.com/zyvd/03/ff/c1f4-dc72-11e9-af2c-00163e042306",
              "urlext": "http://dlvideo.izuiyou.com/zyvd/03/ff/c1f4-dc72-11e9-af2c-00163e042306",
              "urlwm": "http://video.izuiyou.com/zyvd/18/05/1c88-dc72-11e9-af2c-00163e042306",
              "cover_urls": [
                "http://tbfile.izuiyou.com/img/frame/id/825204537?w=540\u0026xcdelogo=0"
              ],
              "type": 0
            }
          },
          "reviews": 79,
          "likes": 519,
          "up": 283,
          "content": "盗墓的生活 就是这么枯燥与乏味#达人冲榜#",
          "c_type": 1
        }
      }
    ],
    "more": 1,
    "offset": 1569069946
  }
}


*/

type FavorDetailReq struct {
	ID     int    `json:"id"`
	HModel string `json:"h_model"`
	HCh    string `json:"h_ch"`
	HApp   string `json:"h_app"`
	CTypes []int  `json:"c_types"`
	HTs    int64  `json:"h_ts"`
	HAv    string `json:"h_av"`
	HNt    int    `json:"h_nt"`
	HDid   string `json:"h_did"`
	Offset int    `json:"offset"`
	HOs    int    `json:"h_os"`
	HM     int    `json:"h_m"`
	Token  string `json:"token"`
	HDt    int    `json:"h_dt"`
}

func GetFavorDetail(favorId int, offset int) ([]*common.VideoModel, error) {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	data := FavorDetailReq{
		ID:     favorId,
		HModel: "iPhone SE",
		HCh:    "appstore",
		HApp:   "zuiyou_speed",
		CTypes: []int{1},
		HTs:    time.Now().UnixNano() / 1000000,
		HAv:    "1.0.8",
		HNt:    1,
		HDid:   "3865a9635d0ed344d4a3b478cbeeae0422f8281e",
		Offset: offset,
		HM:     1298829,
		HOs:    1240000,
		Token:  common.ReadConfig().Zy.Token,
		HDt:    1,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.izuiyou.com/favor/metadata?sign=74cdfa1b2c4e87bee2d47827ded23e9c", body)
	if err != nil {
		return nil, err
	}
	ConfigHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var videoList model.FavorVideoList
	err = json.Unmarshal(buf, &videoList)

	if err != nil {
		return nil, err
	}

	var res []*common.VideoModel
	for _, v := range videoList.Data.List {
		vm := common.VideoModel{}
		vm.Title = TrimTitle(v.Data.Content)
		vm.ID = fmt.Sprintf("%v", v.Data.ID)
		vm.DownLoadDir = DownloadDir()

		for _, video := range v.Data.Videos {
			if vm.Url == "" {
				vm.Url = video.Urlsrc
			}
		}
		vm.SaveKey = SaveKey(vm.Url)
		detail := common.VideoDetail{}
		detail.Tags = append(detail.Tags, v.Data.Topic.Topic)
		detail.Tags = append(detail.Tags, v.Data.Topic.AttsTitle)
		detail.Tags = append(detail.Tags, v.Data.Member.Name)
		if len(v.Data.GodReviews) > 0 {
			for _, g := range v.Data.GodReviews {
				detail.Desc = detail.Desc + "-" + g.Review
			}
		}
		vm.Detail = &detail

		res = append(res, &vm)
	}

	return res, nil

}
