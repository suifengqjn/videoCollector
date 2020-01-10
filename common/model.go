package common

import (
	"encoding/json"
	"fmt"
	"myProject/videoCollector/account"
	"myTool/annie"
	"myTool/file"
	"myTool/ytdl"
	"net/url"
	"os"
	"time"
)

type VideoModel struct {
	Url          string
	DownLoadUrl  string `json:"-"`
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

	u, err := url.Parse(v.Url)
	if err != nil {
		return "", err
	}

	filePath := v.DownLoadDir + "/" + v.Title + ".mp4"

	var isVIP bool
	if account.VcAccount.AccType == 0 {
		isVIP = false
	} else {
		isVIP = true
	}

	if Contains(PxDomains,Domain(u.Host)) {
		if isVIP {
			err = DownLoadWithSSR(v.DownLoadUrl, filePath)
			if err == nil {
				account.VcAccount.DownloadAction()
			}
		} else {
			err = ytdl.DownLoad(v.DownLoadUrl, filePath, ReadConfig().Proxy)
		}
	} else {
		err = annie.DownLoadUrl(v.DownLoadDir, v.Title, v.Url,"")
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

	return ReadConfig().ProjectDir + "/" + time.Now().Format(DownloadTimeFormat) + "/video"

}
