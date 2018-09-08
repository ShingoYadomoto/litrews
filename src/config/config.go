package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

type Conf struct {
	App         *App         `envconfig:"app"`
	Log         *Log         `envconfig:"log"`
	Rss2JsonApi *Rss2JsonApi `envconfig:"rss2jsonapi"`
	LineApi     *LineApi     `envconfig:"lineapi"`
	GoogleApi   *GoogleApi   `envconfig:"googleapi"`
	ViewDir     string
}

type App struct {
	Name   string
	Domain string
	URL    string
}

type Log struct {
	Level log.Lvl
}

type Rss2JsonApi struct {
	Key          string
	BaseEndPoint string
}

type LineApi struct {
	ChannelSecret      string
	ChannelAccessToken string
}

type GoogleApi struct {
	AnalyticsID string
}

func GetConfig() (conf Conf) {
	err := godotenv.Load()
	if err != nil {
		err = errors.Wrap(err, "Error loading .env file")
		log.Error(err)
	}

	err = envconfig.Process("litrews", &conf)
	if err != nil {
		err = errors.Wrap(err, "Error mapping .env file")
		log.Error(err)
	}

	return
}
