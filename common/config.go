package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"time"
)

var conf *GlobalCon

type GlobalCon struct {
	Title       string
	AppID       string     `toml:"appid"`
	ProjectDir  string
	DBFile      string     `toml:"dbFile"`
	TitleLength int      `toml:"title_length"`
	DescLength  int     `toml:"desc_length"`
	Output      string	`toml:"output"`
	Condition   *Condition `toml:"condition"`
	Zy          ZuiYou     `toml:"zy"`
	Youtube     Youtube    `toml:"youtube"`
}

type Condition struct {
	Width      int
	Height     int
	Direction  string
	Size       [2]int
	Similarity float32
	BlackList  []string `toml:"black_list"`
	filter     []string
}

type ZuiYou struct {
	Switch bool
	Filter []string
	UserId int `toml:"user_id"`
	Cookie string
	Token  string
	Favor  struct {
		Group  []string
		Enable bool
		Limit  int
	}
	Attention struct {
		Enable bool
		Limit  int
	}
	Recommend struct {
		Enable bool
		Limit  int
	}
}

type Youtube struct {
	Switch        bool
	Filter        []string
	Keywords      []string
	DurationLimit []int `toml:"duration_limit"`
	TimeLimit     int `toml:"time_limit"`
	Count         int
	Pages []string`toml:"pages"`
}

func ReadConfig() *GlobalCon {
	if conf != nil {
		return conf
	}

	if _, err := toml.DecodeFile("./conf/config.toml", &conf); err != nil {

		path := os.ExpandEnv("$HOME") + "/DeskTop/media/vc_mac/conf/config.toml"
		if _, err := toml.DecodeFile(path, &conf); err != nil {
			fmt.Println("配置文件出错，请检查 config.toml")
			fmt.Println(err)
			time.Sleep(time.Second * 30)
			return nil
		}

	}

	conf.ProjectDir, _ = os.Getwd()

	return conf

}

func ReadDebugConfig() *GlobalCon {
	if conf != nil {
		return conf
	}
	path := os.ExpandEnv("$HOME") + "/go/src/myProject/videoCollector/conf/config.toml"
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		panic(err)
	}
	conf.ProjectDir, _ = os.Getwd()
	return conf

}
