package model

type ZuiyouVideoList struct {
	Ret  int `json:"ret"`
	Data struct {
		List []struct {
			Type int `json:"type"`
			Data struct {
				GodReviews []struct {
					Likes  int    `json:"likes"`
					Up     int    `json:"up"`
					Down   int    `json:"down"`
					Isgod  int    `json:"isgod"`
					Review string `json:"review"`
					Imgs   []struct {
						Urls struct {
							Num360 struct {
								H    int      `json:"h"`
								W    int      `json:"w"`
								Urls []string `json:"urls"`
							} `json:"360"`
							Num540 struct {
								H    int      `json:"h"`
								W    int      `json:"w"`
								Urls []string `json:"urls"`
							} `json:"540"`
						} `json:"urls"`
					} `json:"imgs"`
					Videos struct {
						Num822581410 struct {
							Dur       int      `json:"dur"`
							Thumb     int      `json:"thumb"`
							Playcnt   int      `json:"playcnt"`
							URL       string   `json:"url"`
							Priority  int      `json:"priority"`
							Urlsrc    string   `json:"urlsrc"`
							Urlext    string   `json:"urlext"`
							Urlwm     string   `json:"urlwm"`
							CoverUrls []string `json:"cover_urls"`
							Score     float64  `json:"score"`
							Escore    int      `json:"escore"`
							Type      int      `json:"type"`
						} `json:"822581410"`
					} `json:"videos"`
					Audit   int    `json:"audit"`
					AppName string `json:"app_name"`
				} `json:"god_reviews"`
				Videos struct {
					Num822155452 struct {
						URL       string   `json:"url"`
						Priority  int      `json:"priority"`
						Urlsrc    string   `json:"urlsrc"`
						CoverUrls []string `json:"cover_urls"`
						Type      int      `json:"type"`
					} `json:"822155452"`
				} `json:"videos"`
				Reviews int    `json:"reviews"`
				Likes   int    `json:"likes"`
				Up      int    `json:"up"`
				Down    int    `json:"down"`
				Content string `json:"content"`
				Imgs    []struct {
					Urls struct {
						Num360 struct {
							H    int      `json:"h"`
							W    int      `json:"w"`
							Urls []string `json:"urls"`
						} `json:"360"`
						Num540 struct {
							H    int      `json:"h"`
							W    int      `json:"w"`
							Urls []string `json:"urls"`
						} `json:"540"`
					} `json:"urls"`
				} `json:"imgs"`
			} `json:"data"`
		} `json:"list"`
		More   int `json:"more"`
		Offset int `json:"offset"`
	} `json:"data"`
	More int`json:"more,omitempty"`
	Offset int`json:"offset,omitempty"`
}



