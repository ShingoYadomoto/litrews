package api

import (
	"html/template"
	"strings"

	"github.com/ShingoYadomoto/litrews/src/model"
)

type RssData struct {
	RssURL      string    `json:"rss"`
	Link        string    `json:"link"`
	Count       int       `json:"num"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EndTime     int       `json:"endTime"`
	StartTime   int       `json:"startTime"`
	PublishDate int       `json:"pubDate"`
	ProcessTime int       `json:"processTime"`
	Log         string    `json:"log"`
	Articles    []Article `json:"entries"`
}

type Article struct {
	Link        string `json:"link"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishDate int    `json:"pubDate"`
	Body        template.HTML
}

// "https://news.google.com/news/rss/headlines/section/topic/{TOPIC}.ja_jp/{トピック}?ned=jp&hl=ja&gl=JP" 形式のgoogleNews rssエンドポイントを返す
func GetGoogleNewsEndPointByTopic(topic *model.Topic) (endPoint string) {
	endPoint = "https://news.google.com/news/rss/headlines/section/topic/"
	params := "?ned=jp&hl=ja&gl=JP"
	endPoint += strings.ToUpper(topic.Name) + ".ja_jp/" + topic.NameJa + params

	return
}
