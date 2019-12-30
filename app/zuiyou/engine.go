package zuiyou

import (
	"myProject/videoCollector/app/zuiyou/api"
	"myProject/videoCollector/collector"
	"myProject/videoCollector/common"
)

type Engine struct {

	conf *common.GlobalCon
}

func NewEngine(conf *common.GlobalCon) *Engine  {
	return &Engine{conf:conf}
}

func (e *Engine)Fetch(collector *collector.Collector)  {

	if !e.conf.Zy.Switch {
		return
	}

	if e.conf.Zy.Favor.Enable {

		res := api.GetFavorList()

		collector.PushVideos(res)

	}

	if e.conf.Zy.Recommend.Enable {

		res := api.GetRecommend()
		collector.PushVideos(res)

	}

	if e.conf.Zy.Attention.Enable {

		res := api.GetAttentionUp()

		collector.PushVideos(res)

	}



}


func (e *Engine)Identity()string  {
	return "zuiyou"
}