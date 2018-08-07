package config

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	App *App `yaml:"App"`
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
