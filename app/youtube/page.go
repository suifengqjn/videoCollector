package youtube

import (
	"fmt"
	"github.com/Aiicy/htmlquery"
	html2 "golang.org/x/net/html"
	"myProject/videoCollector/collector"
	"myProject/videoCollector/common"
	util2 "myTool/util"
	"myTool/xpath"
	"time"
)

func (e *Engine) FetchPageVideos(pages []string, collector *collector.Collector) {

	if len(pages) == 0 {
		return
	}

	go func() {

		for {
			IDs := <-e.channel

			if e.CanUse() == false {
				fmt.Println("今日下载次数已用完，请明日再试！")
				time.Sleep(time.Hour)
				break
			}
			videos := e.getVideosByIds(IDs)
			for _, v := range videos {
				v.Pass = true
			}
			if len(videos) > 0 {
				collector.PushVideos(videos)
			}

		}

	}()

	for _, p := range pages {
		e.fetchPage(p, e.channel)
	}
}
func (e *Engine) fetchPage(pageUrl string, channel chan []string) {
	time.Sleep(time.Second)
	var top *html2.Node
	var err error
	var videoIds []string
	top, err = xpath.FetchWithClient(pageUrl, e.client.GetClient(), common.BrowserHeader())
	if err != nil || top == nil{
		e.client.Update()
		return
	}

	html := htmlquery.OutputHTML(top, true)

	arr := rex.FindAllStringSubmatch(html, -1)

	for _, a := range arr {
		if len(a) < 2 || len(a[1]) < 8 {
			continue
		}
		videoIds = append(videoIds, a[1])
	}
	videoIds = util2.RemoveDuplicateElement(videoIds)
	channel <- videoIds

}
