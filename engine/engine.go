package engine

import (
	"myProject/videoCollector/account"
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
	Account    *account.Account
	Apps      []Fetcher
	Collector *collector.Collector
	conf      *common.GlobalCon
}

func NewEngine(conf *common.GlobalCon) *Engine {

	collector := collector.NewCollector()
	zy := zy.NewEngine(conf)
	yt := yt.NewEngine(conf, account.VcAccount)
	apps := []Fetcher{zy, yt}

	return &Engine{
		Account:account.VcAccount,
		Apps:      apps,
		Collector: collector,
		conf:      conf,
	}
}

func (e *Engine) Init() {

}


func (e *Engine) Run() {

	go e.Collector.Run()
	e.work()
	if e.conf.Task > 0 {
		ticker := time.NewTicker(time.Hour * time.Duration(e.conf.Task))
		for range ticker.C {
			e.work()
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

