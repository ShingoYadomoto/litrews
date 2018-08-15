package api

import (
	"html/template"
	"strings"
)

var GoogleNewsTopic = map[string]string{
	"nation":        "日本",
	"world":         "国際",
	"business":      "ビジネス",
	"politics":      "政治",
	"entertainment": "エンタメ",
	"sports":        "スポーツ",
	"sciTech":       "科学&テクノロジー",
}

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
func GetGoogleNewsEndPoint(topic string) (endPoint string) {
	endPoint = "https://news.google.com/news/rss/headlines/section/topic/"
	params := "?ned=jp&hl=ja&gl=JP"
	endPoint += strings.ToUpper(topic) + ".ja_jp/" + GoogleNewsTopic[topic] + params

	return
}
