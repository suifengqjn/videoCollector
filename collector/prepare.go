package collector

import (
	"myProject/videoCollector/common"
	"myProject/videoCollector/library/filter/toutiao"
)

//判断是否已经下载过，过滤平台

func (c *Collector) check(video common.VideoModel) bool   {


	down := c.hasDownLoad(video)
	if down {
		return false
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