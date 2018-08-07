package config

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	App      *App     `yaml:"App"`
	Log      *Log     `yaml:"Log"`
	Database Database `yaml:"Database"`
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
