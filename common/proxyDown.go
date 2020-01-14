package common

import (
	"myTool/proxyClient"
	"myTool/ytdl"
	"net/http"
	"time"
)

var Client *ClientManager

type ClientManager struct {
	Target string
	Client *proxyClient.ProxyClient
}

func NewClientManager() *ClientManager {
	target := "https://www.youtube.com"
	proC, err := proxyClient.NewProxyClient(target)
	if err != nil {
		panic("网络不可用")
	}
	Client = &ClientManager{target, proC}
	return Client
}

func (c *ClientManager)GetClient()*http.Client  {
	return c.Client.MakeClient(time.Second * 30)
}

func (c *ClientManager)GetDownLoadClient()*http.Client  {
	return c.Client.MakeDownLoadClient()
}

func (c *ClientManager)Update()  {
	c.Client.Update()
}

func DownLoadWithSSR(url, path string) error  {
	cli := Client.GetDownLoadClient()
	if cli == nil {
		return ProxyError
	}
	err := ytdl.DownLoadWithClient(url, path, cli)
	return err
}
//
//
//const target = "https://www.youtube.com"
//var accounts []string
//
//var CurrentSSR string
//func LoadSSRAccounts() []string {
//
//	if len(accounts) > 0 {
//		return accounts
//	}
//	str := client.GetAPISSRAccount()
//	decodeBytes, err := base64.StdEncoding.DecodeString(str)
//	if err != nil {
//		return nil
//	}
//
//	accounts= strings.Split(string(decodeBytes), "\n")
//	return accounts
//}
//
//func GetDownLoadClient() *http.Client  {
//
//	return check.MakeDownloadClient(CurrentSSR)
//
//}
//
//func GetClient() *http.Client  {
//
//	return check.MakeClient(CurrentSSR, time.Second * 10)
//
//}
//
//func NewSSR()  {
//	accs := LoadSSRAccounts()
//	var cli *http.Client
//	for _, a := range accs {
//		cli = check.CheckClient(a, target)
//		if cli != nil {
//
//			CurrentSSR = a
//			break
//		}
//	}
//}
//
//func DownLoadWithSSR(url, path string) error  {
//	cli := GetDownLoadClient()
//	if cli == nil {
//		return ProxyError
//	}
//	err := ytdl.DownLoadWithClient(url, path, cli)
//	return err
//}
