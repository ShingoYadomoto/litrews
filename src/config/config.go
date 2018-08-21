package config

import (
	"io/ioutil"

	"github.com/labstack/gommon/log"
	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	App         *App        `yaml:"App"`
	Log         *Log        `yaml:"Log"`
	Database    Database    `yaml:"Database"`
	DocomoApi   DocomoApi   `yaml:"DocomoApi"`
	Rss2JsonApi Rss2JsonApi `yaml:"Rss2JsonApi"`
	LineApi     LineApi     `yaml:"LineApi"`
}

type App struct {
	Name   string `yaml:"Name"`
	Domain string `yaml:"Domain"`
	Port   int    `yaml:"Port"`
}

type Log struct {
	Level log.Lvl `yaml:"Level"`
}

type Database struct {
	Name                  string `yaml:"Name"`
	Dialect               string `yaml:"Dialect"`
	Role                  string `yaml:"Role"`
	Addr                  string `yaml:"Addr"`
	DBName                string `yaml:"DBName"`
	User                  string `yaml:"User"`
	Password              string `yaml:"Password"`
	Net                   string `yaml:"Net"`
	MaxConnections        int    `yaml:"MaxConnections"`
	MaxIdleConnections    int    `yaml:"MaxIdleConnections"`
	ConnectionMaxLifeTime int    `yaml:"ConnectionMaxLifeTime"` // seconds
	Logging               bool   `yaml:"Logging"`               // true or false
}

type DocomoApi struct {
	Key          string `yaml:"Key"`
	BaseEndPoint string `yaml:"BaseEndPoint"`
	CommonParams string `yaml:"CommonParams"`
}

type Rss2JsonApi struct {
	Key          string `yaml:"Key"`
	BaseEndPoint string `yaml:"BaseEndPoint"`
}

type LineApi struct {
	ChannelSecret      string `yaml:"ChannelSecret"`
	ChannelAccessToken string `yaml:"ChannelAccessToken"`
}

func Load(path string) *Conf {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	var conf Conf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		log.Panic(err)
	}
	return &conf
}
