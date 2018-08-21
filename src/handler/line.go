package handler

import (
	"fmt"

	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Submit(c echo.Context) (err error) {
	cc := c.(*context.CustomContext)

	conf := api.LineApi{cc.GetConfig().LineApi}

	bot, err := linebot.New(conf.ChannelSecret, conf.ChannelAccessToken)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Print(bot)
	return
}
