package youtube

import (
	"myProject/videoCollector/collector"
)

func (e *Engine) FetchUrlVideos(IDs []string, collector *collector.Collector) {

	if len(IDs) == 0 {
		return
	}

	videos := e.getVideosByIds(IDs)
	for _,v := range videos {
		v.Pass = true
	}
	if len(videos) > 0 {
		collector.PushVideos(videos)
	}
}