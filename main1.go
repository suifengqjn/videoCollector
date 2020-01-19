package main

import (
	"fmt"
	"log"
	"myTool/ytdl"
)

func main() {

	p := "127.0.0.1:1086"
	vid, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=d72zyGLnpIA",p)
	if err != nil {
		fmt.Println("Failed to get video info")
		return
	}
	fmt.Println(vid.Duration.Minutes())
	fmt.Println(vid.Keywords)

	path := vid.Title + ".mp4"
	err = ytdl.DownLoad(vid.DownLoadUrl, path, p)
	log.Println(err)
	//file, _ := os.Create(vid.Title + ".mp4")
	//defer file.Close()
	//vid.Download(vid.Formats[0], file)

}