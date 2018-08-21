package handler

import (
	"net/http"

	"html/template"
	"net/url"

	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/labstack/echo"
)

func Home(c echo.Context) (err error) {
	cc := c.(*context.CustomContext)
	conf := cc.GetConfig()

	topics := api.GetTopics()

	googleNewsEndPoint := api.GetGoogleNewsEndPointByTopic(&topics[0])
	encodedGoogleNewsEndPoint := url.QueryEscape(googleNewsEndPoint)

	articlesData := new(api.RssData)

	rss2JsonApi := api.Rss2JsonApi{conf.Rss2JsonApi}
	rss2JsonEndPoint := rss2JsonApi.GetEndPoint(encodedGoogleNewsEndPoint)

	rss2JsonApi.SetEncodedDataFromEndPoint(articlesData, rss2JsonEndPoint)

	for i, article := range articlesData.Articles {
		articlesData.Articles[i].Body = template.HTML(article.Description)
	}

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"articlesData": articlesData,
	})
}
