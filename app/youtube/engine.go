package youtube

import (
	"myProject/videoCollector/collector"
	"myProject/videoCollector/common"


)

type Engine struct {
	conf *common.GlobalCon
	durationLimit float64
	channel chan []string
}

func NewEngine(conf *common.GlobalCon) *Engine  {

	channel := make(chan []string, 200)
	return &Engine{conf:conf,durationLimit:60.0, channel:channel}
}

func (e *Engine)Fetch(collector *collector.Collector)  {


	if e.conf.Youtube.Switch {
		keyWords := e.conf.Youtube.Keywords
		count := e.conf.Youtube.Count

		e.FetchKeywords(keyWords, count, collector)
	}




}

func (e *Engine)Identity()string  {
	return "youtube"
}