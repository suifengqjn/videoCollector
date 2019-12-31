package youtube

import (
	"fmt"
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

	url := fmt.Sprintf("https://www.youtube.com/watch?v=%v", ID)

	info, err := ytdl.GetVideoInfo(url, e.conf.Proxy)
	if err != nil {
		return nil
	}

	// duration limit
	if info.Duration.Minutes() > float64(e.conf.Youtube.DurationLimit) {
		return nil
	}

	detail := common.VideoDetail{
		Tags: info.Keywords,
		Desc: info.Description,
	}
	video := common.VideoModel{
		Url:         url,
		DownLoadUrl: info.DownLoadUrl,
		ID:          ID,
		Title:       info.Title,
		DownLoadDir: DownloadDir(),
		SaveKey:     SaveKey(ID),
		Detail:      &detail,
	}

	return &video

}

func DownloadDir() string {

	dir := "./" + time.Now().Format(common.DownloadTimeFormat)
	if file.PathExist(dir) == false {
		_ = os.Mkdir(dir, os.ModePerm)

	}
	return dir + "/youtube"
}

func SaveKey(ID string) []byte {
	return []byte(fmt.Sprintf("youtube_%v", ID))
}
