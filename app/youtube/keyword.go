package youtube

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	html2 "golang.org/x/net/html"
	"myProject/videoCollector/collector"
	"myProject/videoCollector/common"
	util2 "myTool/common"
	"myTool/xpath"
	"regexp"
	"time"
)

/*
&sp=EgIIAg%253D%253D today
&sp=EgQIAxAB this week
&sp=EgQIBBAB this month
&sp=EgQIBRAB this year
*/

var timeLimit = map[int]string{
	1: "&sp=EgIIAg%253D%253D",
	2: "&sp=EgQIAxAB",
	3: "&sp=EgQIBBAB",
	4: "&sp=EgQIBRAB",
}

/*
1 video
2 playlist
3 movie
4 show
*/
var videoType = map[int]string{
	1: "&sp=EgIQAQ%253D%253D",
	2: "&sp=EgIQAw%253D%253D",
	3: "&sp=EgIQBA%253D%253D",
	4: "&sp=EgIQBQ%253D%253D",
}
/*
1 小于 4 minutes
2 大于 20 minute
*/
var durationLimit = map[int]string {
	1:"&sp=EgIYAQ%253D%253D",
	2:"&sp=EgIYAg%253D%253D",

}

func (e *Engine) getTimeLimit() string {
	if e.conf.Youtube.TimeLimit > 0 {

		if _, ok := timeLimit[e.conf.Youtube.TimeLimit]; ok {
			return timeLimit[e.conf.Youtube.TimeLimit]
		}
	}
	return ""
}

func (e *Engine) FetchKeywords(words []string, count int, collector *collector.Collector) {

	go func() {

		for {
			IDs := <-e.channel

			if e.CanUse() == false {
				fmt.Println("今日下载次数已用完，请明日再试！")
				time.Sleep(time.Hour)
				break
			}
			videos := e.getVideosByIds(IDs)
			if len(videos) > 0{
				collector.PushVideos(videos)
			}

		}

	}()

	for _, w := range words {
		e.fetchOneKeyword(w, count, e.channel)
	}
}

func (e *Engine) fetchOneKeyword(word string, count int, channel chan []string) {

	url := fmt.Sprintf("https://www.youtube.com/results?search_query=%v", word) + e.getTimeLimit()

	if len(e.conf.Youtube.DurationLimit) == 2 {
		url += durationLimit[e.conf.Youtube.DurationLimit[1]]
	}

	e.GetVideoIds(url, count, 1, nil, channel)

}

func (e *Engine) getVideosByIds(IDs []string) []*common.VideoModel {
	var res []*common.VideoModel
	for _, id := range IDs {
		v := e.GetVideoInfo(id)
		if v != nil {
			res = append(res, v)
		}
	}

	return res
}


var rex = regexp.MustCompile(`watch\?v=([a-zA-Z0-9-_]+)(?:"|)?`)

func (e *Engine) GetVideoIds(url string, limit int, index int, videoIds []string, channel chan []string) []string {

	time.Sleep(time.Second)
	if collector.Queue.Queue.Len() > 20 {
		time.Sleep(time.Second * 5)
	}

	if e.CanUse() == false {
		return nil
	}
	var fetchUrl = url
	if len(videoIds) >= int(limit) {
		return videoIds
	} else {
		if index >= 2 {
			fetchUrl = fmt.Sprintf("%v&page=%v", url, index)
		}
	}

	var top *html2.Node
	var err error
	top, err = xpath.FetchWithClient(fetchUrl, e.client.GetClient(), common.BrowserHeader())
	if err != nil || top == nil{
		e.client.Update()
	}

	if err != nil {
		return e.GetVideoIds(url, limit, index+1, videoIds, channel)
	}

	html := htmlquery.OutputHTML(top, true)

	arr := rex.FindAllStringSubmatch(html, -1)
	var temIds []string
	for _, a := range arr {
		if len(a) < 2 || len(a[1]) < 8 {
			continue
		}
		videoIds = append(videoIds, a[1])
		temIds = append(temIds, a[1])
	}
	temIds = util2.RemoveDuplicateElement(temIds)
	if len(temIds) > 0 {
		channel <- temIds
	}

	return e.GetVideoIds(url, limit, index+1, videoIds, channel)
}
