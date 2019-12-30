package engine

import (
	yt "myProject/videoCollector/app/youtube"
	zy "myProject/videoCollector/app/zuiyou"
	"myProject/videoCollector/collector"
	"myProject/videoCollector/common"
	"time"
)

type Fetcher interface {
	Fetch(collector *collector.Collector)
	Identity() string
}

type Engine struct {
	Apps      []Fetcher
	Collector *collector.Collector
	conf      *common.GlobalCon
}

func NewEngine(conf *common.GlobalCon) *Engine {

	collector := collector.NewCollector()
	zy := zy.NewEngine(conf)
	yt := yt.NewEngine(conf)
	apps := []Fetcher{zy, yt}



	return &Engine{
		Apps:      apps,
		Collector: collector,
		conf:      conf,
	}

}

func (e *Engine) Run() {

	go e.work()
	go e.Collector.Run()

	ticker := time.NewTicker(time.Hour)

	for range ticker.C {
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
		go app.Fetch(e.Collector)
	}
}
