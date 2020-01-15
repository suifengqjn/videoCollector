package youtube

import (
	"fmt"
	yd "github.com/rylio/ytdl"
	"log"
	"myProject/videoCollector/common"
	"myTool/file"
	"myTool/ytdl"
	"os"
	"time"
)

func (e *Engine) GetVideoInfo(ID string) *common.VideoModel {

	if err := recover(); err != nil {
		return nil
	}
	time.Sleep(time.Second)
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%v", ID)

	var info *yd.VideoInfo
	var err error
	info, err = ytdl.GetVideoInfoWithClient(url, e.client.GetClient())
	if err != nil {
		e.client.Update()
	}

	if err != nil {
		return nil
	}

	if info.Duration.Minutes() > e.durationLimit {
		return nil
	}

	// duration limit
	if len(e.conf.Youtube.DurationLimit) == 2 {
		min := e.conf.Youtube.DurationLimit[0]
		max := e.conf.Youtube.DurationLimit[1]
		if info.Duration.Minutes() < float64(min) || info.Duration.Minutes() > float64(max) {
			return nil
		}
	}

	title := info.Title
	log.Println(title)
	if len(title) == 0 {
		e.client.Update()
		return nil
	}
	if common.ReadConfig().TitleLength > 0 {

		title = common.ExtractTitle(title,common.ReadConfig().TitleLength)
	}

	desc := info.Description

	if common.ReadConfig().DescLength > 0 {
		desc = common.ExtractDesc(desc,common.ReadConfig().DescLength)
	}

	detail := common.VideoDetail{
		Tags: info.Keywords,
		Desc: desc,
	}
	video := common.VideoModel{
		Url:         url,
		DownLoadUrl: info.DownLoadUrl,
		ID:          ID,
		Title:       title,
		DownLoadDir: DownloadDir(),
		SaveKey:     SaveKey(ID),
		Detail:      &detail,
	}

	return &video

}

func DownloadDir() string {

	dir := common.ReadConfig().Output
	var err error
	if len(dir) > 0 {
		if file.PathExist(dir) == false {
			err = os.MkdirAll(dir, os.ModePerm)
		}
	}
	if err != nil {
		dir = "./" + time.Now().Format(common.DownloadTimeFormat)
		if file.PathExist(dir) == false {
			_ = os.Mkdir(dir, os.ModePerm)

		}
		return dir + "/youtube"
	} else {
		return dir
	}

}

func SaveKey(ID string) []byte {
	return []byte(fmt.Sprintf("youtube_%v", ID))
}
