package common

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

var regTitle = regexp.MustCompile("[,。：:!！?\"《》0-9\u4e00-\u9fa5]")
var regDesc = regexp.MustCompile("[#<>《》,。：:!！?\"【】0-9a-zA-Z\u4e00-\u9fa5]")

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

func ExtractDesc(str string, len int) string {

	if strings.Contains(str,"欢迎订阅") {
		return ""
	}
	// 1.
	arr := regDesc.FindAllString(str, -1)
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

func ChineseLen(str string) int  {
	return utf8.RuneCount([]byte(str))
}