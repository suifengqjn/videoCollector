package common

import (
	"io/ioutil"
	"myTool/proxyClient"
	"myTool/ssrClient/check"
	"myTool/ytdl"
	"net/http"
	"strings"
	"time"
)

var Client *ClientManager

type ClientManager struct {
	Target string
	Client *proxyClient.ProxyClient
}

func NewClientManager(vip bool) *ClientManager {

	local := readLocalSSR()
	var free bool
	if vip {
		free = false
	} else {
		free = true
	}

	target := "https://www.youtube.com"
	proC, err := proxyClient.NewProxyClientFree(target, free)
	if err != nil {
		panic("网络不可用")
	}
	proC.AddAdditional(local)
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

func readLocalSSR() []string  {
	buf, err := ioutil.ReadFile("./conf/ssr.txt")
	if err != nil {
		return nil
	}

	arr := strings.Split(string(buf), "\n")
	if len(arr) > 0 {
		var res []string
		for _, s := range arr {
			if strings.HasPrefix(s, "ssr://") {
				if check.CheckUseful(s) {
					res = append(res, s)
				}
			}
		}
		return res
	} else {
		return nil
	}
}