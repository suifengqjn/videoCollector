package collector

import (
	"fmt"
	"log"
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
		log.Println("pop",video.Title)
		c.DealVideos(*video)

	}

}

func(c *Collector)PushVideos(videos []*common.VideoModel)  {

		for _, v := range videos {

			log.Println("push",v.Title)
			c.Queue.Push(v)
		}
}

func (c *Collector) DealVideos(video common.VideoModel) {

	//download check
	pass := c.prepareCheck(video)
	if !pass {
		return
	}

	// download
	path, err := video.DownLoad()
	if err == nil {
		_=c.saveRecord(video)
	}

	fmt.Println(path)

	//remove check
	c.CheckRemove(path)

}
