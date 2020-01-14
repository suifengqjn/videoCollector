package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

var conf *GlobalCon

type GlobalCon struct {
	Title       string
	AppID       string     `toml:"appid"`
	ProjectDir  string
	DBFile      string     `toml:"dbFile"`
	Proxy       string     `toml:"proxy"`
	TitleLength int     `toml:"title_length"`
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
	DurationLimit int `toml:"duration_limit"`
	TimeLimit     int `toml:"time_limit"`
	Count         int
}

func ReadConfig() *GlobalCon {
	if conf != nil {
		return conf
	}

	if _, err := toml.DecodeFile("./conf/config.toml", &conf); err != nil {
		fmt.Println(err)
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
