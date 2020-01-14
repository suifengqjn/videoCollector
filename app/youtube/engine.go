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
	proxy bool
}

func NewEngine(conf *common.GlobalCon,acc *account.Account) *Engine  {

	channel := make(chan []string, 200)
	cli := common.NewClientManager()
	e := &Engine{conf:conf,durationLimit:60.0, channel:channel,client:cli, account:acc}
	e.SetProxy()
	return e
}

func (e *Engine)SetProxy()  {

	if account.VcAccount.AccType > 0 {
		e.proxy = true
	} else {
		e.proxy = false
	}
}

func (e *Engine)Fetch(collector *collector.Collector)  {


	if e.conf.Youtube.Switch {
		keyWords := e.conf.Youtube.Keywords
		count := e.conf.Youtube.Count

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