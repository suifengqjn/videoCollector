package model


type ZyVideo struct {
	Member struct{
		Name string	`json:"name"`
	}	`json:"member"`
	Topic struct {
		ID        int    `json:"id"`
		Topic     string `json:"topic"`
		AttsShow  string `json:"atts_show"`
		AttsTitle string `json:"atts_title"`
		Addition  string `json:"addition"`
		Brief     string `json:"brief"`
		ListShow  string `json:"list_show"`
	} `json:"topic"`
	GodReviews []struct {
		Likes  int    `json:"likes"`
		Up     int    `json:"up"`
		Isgod  int    `json:"isgod"`
		Review string `json:"review"`
		Videos map[string]*VideoDetail `json:"videos"`
	} `json:"god_reviews"`
	PostLabels []struct {
		LabelName string `json:"label_name"`
		LabelID   int    `json:"label_id"`
	} `json:"post_labels"`
	Videos map[string]*VideoDetail `json:"videos"`
	ID      int    `json:"id"`
	Mid     int    `json:"mid"`
	Reviews int64    `json:"reviews"`
	Likes   int    `json:"likes"`
	Up      int    `json:"up"`
	Down    int    `json:"down"`
	Share   int64    `json:"share"`
	Content string `json:"content"`
	Imgs    []struct {
		ID     int    `json:"id"`
		H      int    `json:"h"`
		W      int    `json:"w"`
		Video  int    `json:"video"`
		Dancnt int    `json:"dancnt"`
		Mp4    int    `json:"mp4"`
		Fmt    string `json:"fmt"`
		Urls   map[string]struct {
			H    int      `json:"h"`
			W    int      `json:"w"`
			Urls []string `json:"urls"`
		}  `json:"urls"`
	} `json:"imgs"`
}