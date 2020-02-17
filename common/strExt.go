package common

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

var regTitle = regexp.MustCompile("[-A-Za-z,。：:!！?\"《》0-9\u4e00-\u9fa5]")
var regDesc = regexp.MustCompile("[#<>《》,。：:!！?\"【】0-9\u4e00-\u9fa5]")
var regUrl = regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)

var filterWords = []string{"免費訂閱","欢迎订阅","免费订阅","订阅"}

func ExtractTitle(str string, len int) string {

	// 1.
	arr := regTitle.FindAllString(str, -1)
	str = strings.Join(arr,"")

	// 2.
	if ChineseLen(str) <= len {
		return str
	}

	//3.
	res := ""
	for _, s := range str {
		res = fmt.Sprintf("%v%v", res, string(s))
		if ChineseLen(res) >= len {
			break
		}
	}

	return res
	
}

func ExtractDesc(str string, l int) string {
	for _, w := range filterWords {
		if strings.Contains(str,w) {
			return ""
		}
	}

	// 1. 去掉网址
	urls := regUrl.FindAllString(str,-1)
	if len(urls) > 0 {
		for _,u := range urls {
			str = strings.ReplaceAll(str, u,"")
		}
	}


	// 2. 去掉表情等特殊符号
	arr := regDesc.FindAllString(str, -1)
	str = strings.Join(arr,"")

	// 2.
	if ChineseLen(str) <= l {
		return str
	}

	//3.
	res := ""
	for _, s := range str {
		res = fmt.Sprintf("%v%v", res, string(s))
		if ChineseLen(res) >= l {
			break
		}
	}

	return res

}

func ChineseLen(str string) int  {
	return utf8.RuneCount([]byte(str))
}