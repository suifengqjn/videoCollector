package zuiyou

import (
	"myProject/videoCollector/app/zuiyou/api"
	"myProject/videoCollector/common"
	"myTool/dataStruct/queue"
)

type Engine struct {

	conf *common.GlobalCon
}

func NewEngine(conf *common.GlobalCon) *Engine  {
	return &Engine{conf:conf}
}

func (e *Engine)Fetch(queue *queue.Queue)  {

	if !e.conf.Zy.Switch {
		return
	}

	if e.conf.Zy.Favor.Enable {

		res := api.GetFavorList()

		for _, r := range res {
			queue.Push(r)
		}

	}

	if e.conf.Zy.Recommend.Enable {

		res := api.GetRecommend()

		for _, r := range res {
			queue.Push(r)
		}

	}

	if e.conf.Zy.Attention.Enable {

		res := api.GetAttentionUp()

		for _, r := range res {
			queue.Push(r)
		}

	}



}


func (e *Engine)Identity()string  {
	return "zuiyou"
}