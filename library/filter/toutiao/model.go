package toutiao

type TTSearch struct {
	Count       int      `json:"count"`
	ReturnCount int      `json:"return_count"`
	QueryID     string   `json:"query_id"`
	HasMore     int      `json:"has_more"`
	RequestID   string   `json:"request_id"`
	SearchID    string   `json:"search_id"`
	CurTs       int      `json:"cur_ts"`
	Offset      int      `json:"offset"`
	Message     string   `json:"message"`
	Pd          string   `json:"pd"`
	ShowTabs    int      `json:"show_tabs"`
	Keyword     string   `json:"keyword"`
	City        string   `json:"city"`
	Tokens      []string `json:"tokens"`
	LogPb       struct {
		ImprID string `json:"impr_id"`
	} `json:"log_pb"`
	Data []struct {
		AlaSrc  string `json:"ala_src"`
		AppInfo struct {
			QueryType string `json:"query_type"`
		} `json:"app_info"`
		CellType       int    `json:"cell_type"`
		DisableUa      bool   `json:"disable_ua"`
		DisplayTypeExt string `json:"display_type_ext"`
		Empha          struct {
		} `json:"empha"`
		Host   string   `json:"host"`
		IDStr  string   `json:"id_str"`
		Sign   string   `json:"sign"`
		Title  string   `json:"title"`
		Tokens []string `json:"tokens"`
	} `json:"data"`
	DataHead   []interface{} `json:"data_head"`
	AbFields   interface{}   `json:"ab_fields"`
	Latency    int           `json:"latency"`
	SearchType int           `json:"search_type"`
	TabRank    []string      `json:"tab_rank"`
}
