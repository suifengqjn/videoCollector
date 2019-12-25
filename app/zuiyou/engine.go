package zuiyou

import (
	"myProject/videoCollector/commom"
	"myTool/dataStruct/queue"
	"time"
)

type Engine struct {

	conf *commom.GlobalCon
}

func NewEngine(conf *commom.GlobalCon) *Engine  {
	return &Engine{conf:conf}
}

func (e *Engine)Fetch(queue *queue.Queue)  {


	for {
		time.Sleep(time.Second)

		queue.Push(1)
	}

}


func (e *Engine)Identity()string  {
	return "zuiyou"
}