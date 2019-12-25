package main

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"

	. "github.com/kkdai/youtube"
)

func main() {

	usr, _ := user.Current()
	currentDir := fmt.Sprintf("%v/Movies/youtubedr", usr.HomeDir)
	log.Println("download to dir=", currentDir)
	y := NewYoutube(true)
	arg := "https://www.youtube.com/watch?v=aJOTlE1K90k"
	if err := y.DecodeURL(arg); err != nil {
		fmt.Println("err:", err)
	}
	if err := y.StartDownload(filepath.Join(currentDir, "dl.mp4")); err != nil {
		fmt.Println("err:", err)
	}
}
