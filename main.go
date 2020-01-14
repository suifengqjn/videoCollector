package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	account2 "myProject/videoCollector/account"
	"myProject/videoCollector/common"
	"myProject/videoCollector/engine"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var line = "******************************************************************************"
func main() {


	conf := common.ReadConfig()
	if conf == nil {
		time.Sleep(time.Second * 100)
		fmt.Println("配置文件出错")
		return
	}

	if !check() {
		return
	}

	account := account2.GetAccount(conf.AppID)

	fmt.Println(line)
	fmt.Println(line)
	fmt.Println(line)
	fmt.Println()
	fmt.Println()
	if account.AccType > 0 {
		fmt.Println(formatline("账户ID:"+account.AppId))
		fmt.Println(formatline(account.Msg))
		fmt.Println(formatline(account.Time))
	} else {
		fmt.Println(formatline("免费账户：需要自己配置代理, vip可直接下载"))
		fmt.Println(formatline("vip 购买地址："+"https://www.kuaifaka.com/purchasing?link=3ZUpQ"))
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(line)
	fmt.Println(line)
	fmt.Println(line)


	fmt.Println("开始采集...")

	eng := engine.NewEngine(conf)
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

}

func formatline(text string)string  {

	r := strings.Repeat(" ", 10)
	return r + text + r

}

type Data []struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	HTMLURL       string `json:"html_url"`
	ID            int    `json:"id"`
	NodeID        string `json:"node_id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Labels            []interface{} `json:"labels"`
	State             string        `json:"state"`
	Locked            bool          `json:"locked"`
	Assignee          interface{}   `json:"assignee"`
	Assignees         []interface{} `json:"assignees"`
	Milestone         interface{}   `json:"milestone"`
	Comments          int           `json:"comments"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	ClosedAt          interface{}   `json:"closed_at"`
	AuthorAssociation string        `json:"author_association"`
	Body              string        `json:"body"`
}

type Message struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func check() bool {

	url := "https://api.github.com/repos/suifengqjn/videoCollector/issues"
	client := http.Client{Timeout: time.Second * 20}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}

	var res Data
	err = json.Unmarshal(buf, &res)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}
	var msg Message
	for _, d := range res {
		if d.Title == "1.0" {
			err = json.Unmarshal([]byte(d.Body), &msg)
			break
		}
	}

	if len(msg.Msg) > 0 {
		fmt.Println("===========================")
		fmt.Println(msg.Msg)
		fmt.Println("===========================")
	}

	if msg.Code == 1 {
		return true
	}

	return false
}
