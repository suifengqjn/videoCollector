package common

import (
	"encoding/json"
	"fmt"
	"myTool/file"
	"net/url"
	"os"
	"time"
)

type VideoModel struct {
	Url          string
	DownLoadUrl  string `json:"-"`
	DownLoadUrls []string `json:"-"`
	ID           string
	Title        string
	Detail       *VideoDetail
	PlayCount    int64
	CommentCount int64
	ShareCount   int64
	subVideos    []string
	Up           int
	Down         int
	DownLoadDir  string
	SaveKey      []byte `json:"-"`
	Pic          string
	Platform     string `json:"-"`
	Pass         bool `json:"-"`
}

type VideoDetail struct {
	Desc string
	Tags []string
}

func (v *VideoModel) DownLoad() (string, error) {

	// -o 目录
	// -O 文件名

	if v.DownLoadDir == "" {
		v.DownLoadDir = GetDefaultDownDir()
	}

	if file.PathExist(v.DownLoadDir) == false {
		_ = os.MkdirAll(v.DownLoadDir, os.ModePerm)
	}

	_, err := url.Parse(v.Url)
	if err != nil {
		return "", err
	}

	filePath := v.DownLoadDir + "/" + v.Title + ".mp4"

	fmt.Println("downloading", v.ID)
	err = DownLoadWithSSR(v.DownLoadUrl, filePath)
	if err != nil && len(v.DownLoadUrls) > 0 {
		err = DownLoadWithSSR(v.DownLoadUrls[0], filePath)
	}
	if err != nil {
		_ = os.Remove(filePath)
	} else {
		v.writeToFile()
	}

	return filePath, err
}

func (v *VideoModel) writeToFile() {

	filePath := v.DownLoadDir + "/" + v.Title + ".txt"

	f, _ := os.Create(filePath)

	buf, err := json.Marshal(&v)
	if err == nil {
		_, _ = f.Write(buf)
	}

}

func (v *VideoModel) String() string {
	var tags string
	if v.Detail != nil || len(v.Detail.Tags) > 0 {
		for _, t := range v.Detail.Tags {
			tags = tags + t + "  "
		}
	}

	return fmt.Sprintf("ID = %v \n", v.ID) +
		fmt.Sprintf("title = %v \n", v.Title) +
		fmt.Sprintf("desc = %v \n", v.Detail.Desc) +
		fmt.Sprintf("tags = %v \n", tags)

}

func GetDefaultDownDir() string {

	return ReadConfig().ProjectDir + "/" + time.Now().Format(DownloadTimeFormat)

}
