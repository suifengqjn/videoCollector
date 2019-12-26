package api

import (
	"fmt"
	"testing"
)


func TestGetFavorList(t *testing.T) {

	res := GetFavorList()
	fmt.Println(res)
}

func TestGetFavorDetail(t *testing.T) {
	GetFavorDetail(14496105)
}

func TestGetAttentionUp(t *testing.T) {
	res := GetAttentionUp()
	println(res)
}

func TestGetRecommend(t *testing.T) {
	res := GetRecommend()
	for _, r := range res {
		fmt.Println(r.Title)
	}
}