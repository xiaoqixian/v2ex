// Date:   Wed Jun 11 17:07:03 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package conf

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	JWT struct {
		Secret     string `yaml:"secret"`
		RefExpTime int    `yaml:"refexptime"`
		AccExpTime int    `yaml:"accexptime"`
	} `yaml:"jwt"`

	Rpc struct {
		RpcTimeout int `yaml:"rpctimeout"`
	} `yaml:"rpc"`

	Consul struct {
		User    string `yaml:"user"`
		Post    string `yaml:"post"`
		Comment string `yaml:"comment"`
		Rec     string `yaml:"rec"`
	} `yaml:"consul"`
}

var(
	conf Config
	once sync.Once
)

func Init() {
	f, err := os.ReadFile("conf/conf.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		panic(err)
	}
}

func GetConf() *Config {
	once.Do(Init)
	return &conf
}
