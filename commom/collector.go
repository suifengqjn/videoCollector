package commom

import (
	"fmt"
	"myTool/dataStruct/queue"
	"time"
)
type Collector struct {
	Queue *queue.Queue
}

func NewCollector() *Collector  {
	return &Collector{Queue:queue.NewQueue()}
}

func (c *Collector)Run() {

	for {

		if c.Queue.IsEmpty() {
			time.Sleep(time.Second)
			continue
		}

		v := c.Queue.Pop()

		fmt.Println(v)
		time.Sleep(time.Second * 2)

	}


}