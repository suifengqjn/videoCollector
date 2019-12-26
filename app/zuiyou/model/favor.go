package model

type FavorVideoList struct {
	Ret  int `json:"ret"`
	Data struct {
		List []struct {
			Type int `json:"type"`
			Data struct {
				Member struct{
					Name string	`json:"name"`
				}	`json:"member"`
				Topic struct {
					ID        int    `json:"id"`
					Topic     string `json:"topic"`
					AttsTitle string `json:"atts_title"`
				} `json:"topic"`
				GodReviews []struct {
					Likes  int    `json:"likes"`
					Up     int    `json:"up"`
					Isgod  int    `json:"isgod"`
					Review string `json:"review"`
					Videos map[string]*VideoDetail `json:"videos"`
				} `json:"god_reviews"`
				Videos map[string]*VideoDetail `json:"videos"`
				Reviews int    `json:"reviews"`
				Likes   int    `json:"likes"`
				Up      int    `json:"up"`
				ID      int		`json:"id"`
				Content string `json:"content"`
				CType   int    `json:"c_type"`
			} `json:"data"`
		} `json:"list"`
		More   int `json:"more"`
		Offset int `json:"offset"`
	} `json:"data"`
}

type VideoDetail struct {
	Dur       int      `json:"dur"`
	Thumb     int      `json:"thumb"`
	Playcnt   int      `json:"playcnt"`
	URL       string   `json:"url"`
	Priority  int      `json:"priority"`
	Urlsrc    string   `json:"urlsrc"`
	Urlext    string   `json:"urlext"`
	Urlwm     string   `json:"urlwm"`
	CoverUrls []string `json:"cover_urls"`
	Type      int      `json:"type"`
}