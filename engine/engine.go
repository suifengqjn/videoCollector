package engine

import (
	yt "myProject/videoCollector/app/youtube"
	zy "myProject/videoCollector/app/zuiyou"
	"myProject/videoCollector/commom"
	"myTool/dataStruct/queue"
	"time"
)

type Fetcher interface {
	Fetch(queue *queue.Queue)
	Identity() string
}

type Engine struct {
	Apps      []Fetcher
	Collector *commom.Collector
	conf      *commom.GlobalCon
}

func NewEngine(conf *commom.GlobalCon) *Engine {

	zy := zy.NewEngine(conf)
	yt := yt.NewEngine(conf)
	apps := []Fetcher{zy, yt}

	collector := commom.NewCollector()


	return &Engine{
		Apps:      apps,
		Collector: collector,
		conf:      conf,
	}

}

func (e *Engine) Run() {

	go e.work()
	go e.Collector.Run()

	c := time.Tick(time.Hour)

	for range c {
		h := time.Now().Hour()
		if h >= 9 && h <= 21 {
			e.work()

		} else {
			time.Sleep(time.Hour)
		}
	}

}

func (e *Engine) Stop() {

}

func (e *Engine) work() {
	for _, app := range e.Apps {
		go app.Fetch(e.Collector.Queue)
	}
}
