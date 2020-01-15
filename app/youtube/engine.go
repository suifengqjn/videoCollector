package youtube

import (
	"myProject/videoCollector/account"
	"myProject/videoCollector/collector"
	"myProject/videoCollector/common"


)

type Engine struct {
	conf *common.GlobalCon
	durationLimit float64
	channel chan []string
	client *common.ClientManager
	account *account.Account
}

func NewEngine(conf *common.GlobalCon,acc *account.Account) *Engine  {

	channel := make(chan []string, 200)
	cli := common.NewClientManager(acc.AccType > 0)
	e := &Engine{conf:conf,durationLimit:60.0, channel:channel,client:cli, account:acc}
	return e
}

func (e *Engine)Fetch(collector *collector.Collector)  {


	if e.conf.Youtube.Switch {
		keyWords := e.conf.Youtube.Keywords
		count := e.conf.Youtube.Count

		e.FetchPageVideos(e.conf.Youtube.Pages, collector)
		e.FetchKeywords(keyWords, count, collector)
	}

}

func (e *Engine)CanUse()bool  {

	if e.account.AccType <=0 || e.account.AccType == account.VCVIPSUPERVIP {
		return true
	}
	if e.account.Count <= 0 {
		return false
	}
	return true

}

func (e *Engine)Identity()string  {
	return "youtube"
}