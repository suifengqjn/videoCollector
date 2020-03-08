package youtube

import (
	"myProject/videoCollector/account"
	"myProject/videoCollector/collector"
	"myProject/videoCollector/common"
)

type Engine struct {
	conf          *common.GlobalCon
	durationLimit float64
	channel       chan []string
	client        *common.ClientManager
	account       *account.Account
}

/*
0 基础版月卡  下载速度较慢 设备数量3  ￥5/月
1 多人版月卡  自备ssr账户 无设备数量限制  ￥30/月
2 稳定版月卡  下载速度较快 设备数量3   ￥30/月
3 基础版年卡  下载速度较慢 设备数量5  ￥60/月
4 多人版年卡  自备ssr账户 无设备数量限制  ￥200/月
*/
func NewEngine(conf *common.GlobalCon, acc *account.Account) *Engine {

	channel := make(chan []string, 200)
	isLocal := true
	vip := true
	if acc.AccType == 2 || acc.AccType >= 9 {
		vip = true
	}

	cli := common.NewClientManager(isLocal, vip)
	e := &Engine{conf: conf, durationLimit: 60.0, channel: channel, client: cli, account: acc}
	return e
}

func (e *Engine) Fetch(collector *collector.Collector) {

	if e.conf.Youtube.Switch {
		keyWords := e.conf.Youtube.Keywords
		count := e.conf.Youtube.Count
		go e.FetchUrlVideos(e.conf.Youtube.Urls, collector)
		go e.FetchPageVideos(e.conf.Youtube.Pages, collector)
		e.FetchKeywords(keyWords, count, collector)
	}

}

func (e *Engine) CanUse() bool {

	//if e.account.AccType <=0 || e.account.AccType == account.VCVIPSUPERVIP {
	//	return true
	//}
	//if e.account.Count <= 0 {
	//	return false
	//}
	return true

}

func (e *Engine) Identity() string {
	return "youtube"
}
