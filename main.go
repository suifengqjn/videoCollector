package main

import (
	"fmt"
	"myProject/videoCollector/account"
	"myProject/videoCollector/common"
	"myProject/videoCollector/engine"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var line = "******************************************************************************"
func main() {


	// read config
	conf := common.ReadConfig()
	if conf == nil {
		time.Sleep(time.Second * 100)
		return
	}

	// check version
	//code, msg := account2.CheckVersion()
	//if code != 1 {
	//	if len(msg) > 0 {
	//		fmt.Println(msg)
	//	} else {
	//		fmt.Println("请检查网络，稍后再试")
	//	}
	//	time.Sleep(time.Second * 100)
	//	return
	//} else if len(msg) > 0 {
	//	fmt.Println(msg)
	//	time.Sleep(time.Second * 3)
	//}

	// check account
	appid := account.LoadAppId()
	account := account.NewAccount(appid, "")

	fmt.Println(line)
	fmt.Println(line)
	fmt.Println(line)
	fmt.Println()
	fmt.Println()
	printInfo()
	if account.AccType < 0 {
		fmt.Println(formatline("无效账户"))
		fmt.Println(formatline("密钥 购买地址："+"https://www.kuaifaka.com/purchasing?link=3ZUpQ"))
	} else {
		fmt.Println(formatline(fmt.Sprintf("账户 密钥：%v",account.AppId[:10] + strings.Repeat("*",16))))
		fmt.Println(formatline(account.Msg))
		fmt.Println(formatline(account.Time))
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(line)
	fmt.Println(line)
	fmt.Println(line)

	time.Sleep(time.Second * 6)

	if account.AccType < 0 {
		time.Sleep(time.Second * 100)
		return
	}

	if conf.SSR {
		fmt.Println("ssr账户检测中...")
	}
	eng := engine.NewEngine(conf, account)
	fmt.Println("开始采集...")
	go func() {
		sig := make(chan os.Signal, 1)

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		msg := <-sig

		fmt.Println("receive exit msg:", msg)
		eng.Stop()
		os.Exit(1)
	}()
	eng.Init()
	eng.Run()
	select {}

}

func formatline(text string)string  {

	r := strings.Repeat(" ", 10)
	return r + text + r

}

func printInfo()  {
	fmt.Println(formatline("视频采集器"))
	fmt.Println(formatline("软件地址：https://github.com/suifengqjn/videoCollector"))
}

