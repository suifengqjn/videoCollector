package zuiyou

import (
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



}

func (e *Engine)Identity()string  {
	return "youtube"
}