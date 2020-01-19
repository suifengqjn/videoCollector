package toutiao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myTool/common"
	"myTool/request"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var ttCookie []*http.Cookie

var Enable = false
func init() {
	cur, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	ckPath := fmt.Sprintf("%v/conf/toutiao.json", cur)

	buf, err := ioutil.ReadFile(ckPath)
	if err != nil {
		return
	}

	Enable = true
	ttCookie = request.ParseCookie(buf)
}

//是否已存在类似内容
func CompareTitleExist(title string) bool  {

	searchM := SearchKeyword(title)
	fmt.Println(searchM.Data)
	max := float32(0)
	for _, sea := range searchM.Data {
		v := common.ComparisonString(sea.Title, title)
		if v > max {
			max = v
		}

	}
	fmt.Println("length",len(searchM.Data))
	fmt.Printf("相似度%.2f \n", max)
	if max >= 0.7 {
		return  true
	}

	return false
}



func SearchKeyword(keyWord string) *TTSearch {
	//%E4%B8%AD%E5%9B%BD
	fmt.Printf("toutiao search---keyword :%v" ,keyWord)
	v := url.Values{}
	v.Add("keyword", keyWord)
	kw := v.Encode()
	url := fmt.Sprintf("https://www.toutiao.com/api/search/content/?aid=24&app_name=web_search&offset=0&format=json&%v&autoload=true&count=20&en_qc=1&cur_tab=2&from=video&pd=video&timestamp=%v", kw, time.Now().Unix())
	fmt.Println(url)
	header := GetHeader(kw, url)
	cookie := GetCookies()
	buffer, err :=request.HttpGetBuf(url, nil, header, cookie,"","",0)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer))
	var searchM =  TTSearch{}

	err = json.Unmarshal(buffer, &searchM)

	if err != nil {
		panic(err)
	}

	return &searchM
}


func GetHeader(keyUrl string, url string)map[string]string  {

	//Content-Type	application/json; charset=utf-8
	//               text/plain;charset=UTF-8

	header := make(map[string]string)
	header["X-Requested-With"] = "XMLHttpRequest"
	header["scheme"] = "https"
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"
	header["Content-Type"] = "application/json; charset=utf-8"
	//header["Accept"] = "application/json, text/javascript"
	header["Referer"] = fmt.Sprintf("https://www.toutiao.com/search/?%v", keyUrl)
	//header["Accept-Language"] ="zh-CN,zh;q=0.9,en;q=0.8"
	//header["Accept-Encoding"] ="gzip, deflate, br"
	header["path"] = strings.TrimPrefix(url, "https://www.toutiao.com/api/")
	header["method"] = "GET"
	header["authority"] = "www.toutiao.com"
	header["Connection"]= "keep-alive"

	return header

}


func GetCookies()[]*http.Cookie  {

	return ttCookie
}