package youtube

import (
	"fmt"
	"myProject/videoCollector/account"
	"myProject/videoCollector/common"
	"testing"
)

func TestNewEngine(t *testing.T) {
	// https://www.youtube.com/watch?v=OiBDYWkbIwc
	conf := common.ReadDebugConfig()

	acc := account.GetAccount(conf.AppID)
	fmt.Println(acc)
	eng := NewEngine(conf, acc)

	m := eng.GetVideoInfo("OiBDYWkbIwc")

	fmt.Println(m)

	eng.FetchPageVideos()


}




