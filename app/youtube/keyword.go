package youtube

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"myProject/videoCollector/collector"
	"myProject/videoCollector/common"
	util2 "myTool/util"
	"myTool/xpath"
	"regexp"
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

	var fetchUrl = url
	if len(videoIds) >= int(limit) {
		return videoIds
	} else {
		if index >= 2 {
			fetchUrl = fmt.Sprintf("%v&page=%v", url, index)
		}
	}

	top, err := xpath.FetchHeadCookieProxy(fetchUrl, common.BrowserHeader(), nil, e.conf.Proxy)

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
	channel <- temIds

	return e.GetVideoIds(url, limit, index+1, videoIds, channel)
}
