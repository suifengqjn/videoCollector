package api

import (
	"fmt"
	"myProject/videoCollector/common"
	"testing"
)

func init() {
	_ = common.ReadDebugConfig()
}

func TestGetFavorList(t *testing.T) {

	res := GetFavorList()
	fmt.Println(res)
}

func TestGetFavorDetail(t *testing.T) {
	GetFavorDetail(14496105, 12)
}

func TestGetAttentionUp(t *testing.T) {
	res := GetAttentionUp()
	for _, r := range res {
		fmt.Println(r.Title)
	}
}

func TestGetRecommend(t *testing.T) {
	res := GetRecommend()
	for _, r := range res {
		fmt.Println(r.Title)
	}
}