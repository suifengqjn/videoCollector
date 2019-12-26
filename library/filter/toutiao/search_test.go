package toutiao

import (
	"fmt"
	"myTool/common"
	"testing"
)

func TestSearchKeyword(t *testing.T) {

	v := CompareTitleExist("视频")

	fmt.Println(v)
}

func TestSearchKeyword2(t *testing.T) {

	str1 := "火车压过竟成飞刀，小孩发现了商机"
	str2 := "火车压过竟成飞刀，小孩的so发现了商机!! 我"


	v := common.ComparisonString(str1, str2)

	fmt.Println(v)

}