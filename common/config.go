package common

import "github.com/BurntSushi/toml"

var conf *GlobalCon

type GlobalCon struct {
	Title      string
	ProjectDir string
	DBFile     string    `toml:"dbFile"`
	BlackList  []string  `toml:"black_list"`
	Proxy string	`toml:"proxy"`
	Condition  Condition `toml:"condition"`
	Zy         ZuiYou    `toml:"zy"`
}

type Condition struct {
	Width      int
	Height     int
	Direction  string
	Size       [2]int
	Similarity float32
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

func ReadConfig()*GlobalCon  {
	if conf != nil {
		 return conf
	}

	if _, err := toml.DecodeFile("./conf/config.toml", &conf); err != nil {
		panic(err)
	}

	return conf

}