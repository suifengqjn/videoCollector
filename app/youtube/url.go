package youtube

import (
	"myProject/videoCollector/collector"
)

func (e *Engine) FetchUrlVideos(urls []string, collector *collector.Collector) {

	if len(urls) == 0 {
		return
	}

	videos := e.getVideosByUrls(urls)
	for _,v := range videos {
		v.Pass = true
	}
	if len(videos) > 0 {
		collector.PushVideos(videos)
	}
}