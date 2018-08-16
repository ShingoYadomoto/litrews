package handler

import (
	"html/template"
	"net/http"
	"net/url"

	"strconv"

	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/ShingoYadomoto/litrews/src/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func GoogleNews(c echo.Context) (err error) {
	cc := c.(*context.CustomContext)
	conf := cc.GetConfig()
	db := cc.GetDB()

	topicModel := model.NewTopicModel(db)

	topicID, err := strconv.Atoi(c.Param("topic"))
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	topic, err := topicModel.GetTopicByID(topicID)
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	googleNewsEndPoint := api.GetGoogleNewsEndPointByTopic(topic)
	encodedGoogleNewsEndPoint := url.QueryEscape(googleNewsEndPoint)

	articlesData := new(api.RssData)

	rss2JsonApi := api.Rss2JsonApi{conf.Rss2JsonApi}
	rss2JsonEndPoint := rss2JsonApi.GetEndPoint(encodedGoogleNewsEndPoint)

	rss2JsonApi.SetEncodedDataFromEndPoint(articlesData, rss2JsonEndPoint)

	for i, article := range articlesData.Articles {
		articlesData.Articles[i].Body = template.HTML(article.Description)
	}

	return c.Render(http.StatusOK, "googleNews", map[string]interface{}{
		"articlesData": articlesData,
	})
}
