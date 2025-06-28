// Date:   Wed Jun 11 17:07:03 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package conf

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)


type MySQLConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	DBName    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

type Config struct {
	MySQL MySQLConfig
	Redis RedisConfig
	Kafka KafkaConfig
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
