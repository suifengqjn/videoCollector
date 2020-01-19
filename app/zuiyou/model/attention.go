package model

/*
{
    "ret": 1,
    "data": {
        "list": [
            {
                "topic": {
                    "id": 391280,
                    "topic": "沙雕网友的日常",
                    "atts_show": "7w+沙雕网友",
                    "atts_title": "沙雕网友",
                    "addition": "74411 沙雕网友",
                    "brief": "沙雕又称沙漠之雕，来这里看看来自沙雕网友的沙雕生活吧～",
                    "list_show": "沙雕又称沙漠之雕，来这里看看来自沙雕网友的沙雕生活吧～"
                },
                "god_reviews": [
                    {
                        "mname": "當地梁朝伟",
                        "likes": 208,
                        "up": 85,
                        "isgod": 1,
                        "review": "还在努力",
                        "source": "user",
                        "videos": {
                            "829808736": {
                                "dur": 11,
                                "playcnt": 8341,
                                "url": "http://dlvideo.izuiyou.com/zyvd/f1/03/c21d-e0d2-11e9-af2c-00163e042306",
                                "priority": 1,
                                "urlsrc": "http://dlvideo.izuiyou.com/zyvd/f1/03/c21d-e0d2-11e9-af2c-00163e042306",
                                "urlext": "http://dlvideo.izuiyou.com/zyvd/f1/03/c21d-e0d2-11e9-af2c-00163e042306",
                                "urlwm": "http://dlvideo.izuiyou.com/zyvd/f1/03/c21d-e0d2-11e9-af2c-00163e042306"
                            }
                        }
                    }
                ],
                "post_labels": [
                    {
                        "label_name": "沙雕网友的日常",
                        "label_id": 3111
                    }
                ],
                "videos": {
                    "828640303": {
                        "dur": 10,
                        "url": "http://dlvideo.izuiyou.com/zyvd/1b/54/8e72-df9e-11e9-8ded-00163e02acff",
                        "urlsrc": "http://dlvideo.izuiyou.com/zyvd/1b/54/8e72-df9e-11e9-8ded-00163e02acff",
                        "urlwm": "http://dlvideo.izuiyou.com/zyvd/1c/3f/92f4-df9e-11e9-af2c-00163e042306"
                    }
                },
                "id": 140348182,
                "mid": 209316564,
                "reviews": 108,
                "likes": 658,
                "up": 454,
                "down": 8,
                "share": 11,
                "content": "今天你锻炼了嘛？",
                "imgs": [
                    {
                        "id": 828640303,
                        "h": 540,
                        "w": 540,
                        "video": 1,
                        "dancnt": 30,
                        "mp4": 0,
                        "fmt": "jpeg",
                        "urls": {
                            "540": {
                                "h": 540,
                                "w": 540,
                                "urls": [
                                    "http://tbfile.izuiyou.com/img/frame/id/828640303?w=540&xcdelogo=0"
                                ]
                            }
                        }
                    }
                ]
            }
        ]
    }
}
*/

type AttentionRes struct {
	Ret  int `json:"ret"`
	Data struct {
		List []*ZyVideo `json:"list"`
	} `json:"data"`
}
