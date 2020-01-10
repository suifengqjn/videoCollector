package common

import (
	"encoding/base64"
	"myTool/ssrClient/check"
	"myTool/ssrClient/client"
	"myTool/ytdl"
	"net/http"
	"strings"
	"time"
)

const target = "https://www.youtube.com" 
var accounts []string

var CurrentSSR string
func LoadSSRAccounts() []string {

	if len(accounts) > 0 {
		return accounts
	}
	str := client.GetAPISSRAccount()
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil
	}

	accounts= strings.Split(string(decodeBytes), "\n")
	return accounts
}

func GetDownLoadClient() *http.Client  {

	return check.MakeDownloadClient(CurrentSSR)
	
}

func GetClient() *http.Client  {

	return check.MakeClient(CurrentSSR, time.Second * 10)

}

func NewSSR()  {
	accs := LoadSSRAccounts()
	var cli *http.Client
	for _, a := range accs {
		cli = check.CheckClient(a, target)
		if cli != nil {

			CurrentSSR = a
			break
		}
	}
}

func DownLoadWithSSR(url, path string) error  {
	cli := GetDownLoadClient()
	if cli == nil {
		return ProxyError
	}
	err := ytdl.DownLoadWithClient(url, path, cli)
	return err
}