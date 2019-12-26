package collector

import (
	"fmt"
	"myProject/videoCollector/common"
	"myProject/videoCollector/library/cache/store"
	"myTool/dataStruct/queue"
	"time"
)

type Collector struct {
	Queue *queue.Queue
	store *store.DB
}

func NewCollector() *Collector {

	db, err := store.OpenDB(common.ReadConfig().DBFile)
	if err != nil {
		panic("数据库路径错误")
	}
	return &Collector{Queue: queue.NewQueue(), store: db}
}

func (c *Collector) Run() {

	for {

		if c.Queue.IsEmpty() {
			time.Sleep(time.Second)
			continue
		}

		v := c.Queue.Pop()

		video := v.(*common.VideoModel)
		c.DealVideos(*video)

	}

}

func (c *Collector) DealVideos(video common.VideoModel) {

	//check
	pass := c.check(video)
	if !pass {
		return
	}

	path, err := video.DownLoad()
	if err == nil {
		_=c.saveRecord(video)
	}

	fmt.Println(path)

}
