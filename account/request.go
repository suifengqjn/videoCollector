package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myTool/common"
	"myTool/sign"
	"myTool/sys"
	"net/http"
	"sync"
	"time"
)

const remoteHost = "http://106.12.220.252:8001"

type Account struct {
	AccType int        `json:"acc_type"`
	Count   int        `json:"count"`
	Time    string     `json:"time"`
	Msg     string     `json:"msg"`
	AppId   string     `json:"-"`
	Lock    sync.Mutex `json:"-"`
}

func GetAccountInfo(appId string) *Account {
	url := remoteHost + "/vc/account_info"
	method := "POST"

	url += fmt.Sprintf("?sign=%v", sign.MakeApiSign())

	param := make(map[string]string)
	param["host"] = sys.GetSysInfo().IP
	param["app_id"] = appId

	buf, _ := json.Marshal(param)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(buf))

	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", common.MD5String(fmt.Sprintf("%v", time.Now().UTC().UnixNano())))
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var acc = Account{}

	err = json.Unmarshal(body, &acc)
	if err != nil {
		return nil
	}
	acc.AppId = appId
	acc.Lock = sync.Mutex{}
	return &acc
}

func (a *Account) addCount() error {
	url := remoteHost + "/vc/count"
	url += fmt.Sprintf("?sign=%v", sign.MakeApiSign())
	method := "POST"

	param := make(map[string]string)
	param["app_id"] = a.AppId

	buf, _ := json.Marshal(param)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(buf))

	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", common.MD5String(fmt.Sprintf("%v", time.Now().UTC().UnixNano())))
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if string(body) == "success" {
		return nil
	}
	return err
}

func CheckVersion() (int, string) {
	url := remoteHost + "/vc/check"
	url += fmt.Sprintf("?sign=%v", sign.MakeApiSign())
	url += fmt.Sprintf("&version=%v", Version)
	method := "POST"

	client := &http.Client{}

	var param = make(map[string]string)
	buf, _ := json.Marshal(param)
	req, err := http.NewRequest(method, url, bytes.NewReader(buf))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", common.MD5String(fmt.Sprintf("%v", time.Now().UTC().UnixNano())))
	if err != nil {

		return 0, ""
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	type Message struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	var msg Message
	_ = json.Unmarshal(body, &msg)

	return msg.Code, msg.Msg

}
