package collector

import (
	"myProject/videoCollector/common"
	"myProject/videoCollector/library/filter/toutiao"
	"strings"
)

//判断是否已经下载过，过滤平台

func (c *Collector) prepareCheck(video common.VideoModel) bool   {


	// down
	down := c.hasDownLoad(video)
	if down {
		return false
	}

	if video.Pass {
		return true
	}

	//condition
	condition := common.ReadConfig().Condition

	if len(condition.BlackList) > 0 {
		for _, s := range condition.BlackList {
			if strings.Contains(video.Title, s) {
				return false
			}
		}

	}

	return c.filterPlatform(video)

}

func (c *Collector)hasDownLoad(video common.VideoModel) bool  {

	exist := c.store.Has(video.SaveKey)
	return exist
}

func (c *Collector)saveRecord(video common.VideoModel) error {

	err := c.store.Save(video.SaveKey, []byte("1"))
	return err

}

func (c *Collector)filterPlatform(video common.VideoModel) bool {

	filter := c.getFilter(video.Platform)

	if len(filter) == 0 {
		return true
	}

	for _, f := range filter {
		if f == "t" {
			has := toutiao.CompareTitleExist(video.Title)
			if has  {
				return false
			}

		}
	}

	return true
}

func (c *Collector)getFilter(plat string) []string{
	if plat == common.PlatZY {
		return common.ReadConfig().Zy.Filter
	}
	return nil
}