package api

import (
	"errors"
	"html/template"
	"strings"
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

type Topic struct {
	ID     int
	Name   string
	NameJa string
}

func GetAllTopics() (topics []Topic) {
	topics = []Topic{
		Topic{1, "nation", "日本"},
		Topic{2, "world", "国際"},
		Topic{3, "business", "ビジネス"},
		Topic{4, "politics", "政治"},
		Topic{5, "entertainment", "エンタメ"},
		Topic{6, "sports", "スポーツ"},
		Topic{7, "sciTech", "科学&テクノロジー"},
	}
	return
}

func GetTopicByID(id int) (topic Topic, err error) {
	topics := GetAllTopics()
	for _, v := range topics {
		if v.ID == id {
			topic = v
		}
	}
	if topic.ID == 0 {
		err = errors.New("指定されたトピックは存在しません。")
	}
	return
}

// "https://news.google.com/news/rss/headlines/section/topic/{TOPIC}.ja_jp/{トピック}?ned=jp&hl=ja&gl=JP" 形式のgoogleNews rssエンドポイントを返す
func GetGoogleNewsEndPointByTopic(topic *Topic) (endPoint string) {
	endPoint = "https://news.google.com/news/rss/headlines/section/topic/"
	params := "?ned=jp&hl=ja&gl=JP"
	endPoint += strings.ToUpper(topic.Name) + ".ja_jp/" + topic.NameJa + params

	return
}
