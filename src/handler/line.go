package handler

import (
	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Submit(c echo.Context) (err error) {
	cc := c.(*context.CustomContext)

	conf := api.LineApi{cc.GetConfig().LineApi}

	bot, err := linebot.New(conf.ChannelSecret, conf.ChannelAccessToken)
	_, err := bot.PushMessage(ID, messages...).Do()
	if err != nil {
		// Do something when some bad happened
	}
}
