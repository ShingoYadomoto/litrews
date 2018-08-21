package api

import (
	"github.com/ShingoYadomoto/litrews/src/config"
)

type Rss2JsonApi struct {
	*config.Rss2JsonApi
}

// rssのエンドポイントを渡すと、RSS2JsonAPIのエンドポイントを返す
func (rss *Rss2JsonApi) GetEndPoint(rssEndPoint string) (ep string) {
	ep = rss.BaseEndPoint + "?access_token=" + rss.Key + "&rss=" + rssEndPoint
	return
}

func (rss *Rss2JsonApi) SetEncodedDataFromEndPoint(i interface{}, ep string) (err error) {
	err = setEncodedDataFromEndPoint(i, ep)
	return
}
