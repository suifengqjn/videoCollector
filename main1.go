package main

import (
	"fmt"
	"github.com/rylio/ytdl"
)

func main() {
	ytdl.Socket5Proxy = "127.0.0.1:1086"
	vid, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=EhPA-t_VewU")
	if err != nil {
		fmt.Println("Failed to get video info")
		return
	}
	fmt.Println(vid.Duration.Minutes())
	fmt.Println(vid.Keywords)


	//file, _ := os.Create(vid.Title + ".mp4")
	//defer file.Close()
	//vid.Download(vid.Formats[0], file)

}