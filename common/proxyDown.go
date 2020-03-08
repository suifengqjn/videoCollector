package common

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"myTool/proxyClient"
	"myTool/ssrClient/check"
	"myTool/ytdl"
	"net/http"
	"os"
	"strings"
	"time"
)

var Client *ClientManager

type ClientManager struct {
	Target string
	Client proxyClient.ProxyClientInter
}

func NewClientManager(isLocal,vip bool) *ClientManager {
	target := "https://www.youtube.com"
	local := readLocalSSR()
	fmt.Printf("一共读取ssr账户 %v 个, 正在提取有效ssr \n", len(local))
	if isLocal {
		if len(local) == 0 {
			fmt.Println("SSR 账户无效，请在 conf/ssr.txt 文件中写入自己的SSR账户")
			time.Sleep(time.Hour)
			os.Exit(1)
		}
		proC, _ := proxyClient.NewProxyLocalClient(target,local)
		if len(proC.Accounts) == 0 {
			fmt.Println( " 没有有效 ssr,无法采集")
		}
		Client = &ClientManager{target, proC}
	} else {
		var free bool
		if vip {
			free = false
		} else {
			free = true
		}

		proC, err := proxyClient.NewProxyClientFree(target, free)
		if err != nil {
			fmt.Println("网络异常，请检查网络后再次尝试")
			time.Sleep(time.Hour)
			os.Exit(1)
		}
		proC.AddAdditional(local)
		Client = &ClientManager{target, proC}
	}


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

// 支持单个ssr 和 订阅地址
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
			} else if strings.HasPrefix(s,"http") {
				res = append(res, getSubSSR(s)...)
			}
		}
		return res
	} else {
		return nil
	}
}

func getSubSSR(url string) []string  {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil
	}

	decodeBytes, err := base64.StdEncoding.DecodeString(string(buf))

	arr := strings.Split(string(decodeBytes), "\n")

	var resArr []string
	for _, s := range arr {
		if check.CheckUseful(s) {
			resArr = append(resArr, s)
		}
	}
	return resArr
}